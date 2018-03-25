package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	channelId  = flag.String("channel-id", "UCLRCoLpDImkbFw8D3qXFpUQ", "Channel Id exemple")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

type Configuration struct {
	developerKey string
}

type video struct {
	title       string
	description string
}

func main() {
	videos := VideoList(*channelId)
	printIDs("Videos", videos)
}

func VideoList(channelId string) map[string]video {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("key")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Activities.List("id,snippet").
		ChannelId(channelId).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	videos := make(map[string]video)

	for _, item := range response.Items {
		videos[item.Id] = video{title: item.Snippet.Title, description: item.Snippet.Description}
	}

	return videos
}

func printIDs(sectionName string, matches map[string]video) {
	fmt.Printf("%v:\n", sectionName)
	for id, video := range matches {
		fmt.Printf("[%v] %v\n", id, video.title)
	}
	fmt.Printf("\n\n")
}
