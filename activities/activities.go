package activities

import (
	"flag"
	"fmt"
	"log"

	youtube "google.golang.org/api/youtube/v3"
)

var maxResults = flag.Int64("max-results", 25, "Max YouTube results")

type video struct {
	title        string
	description  string
	thumbnailURL string
	channelTitle string
	channelID    string
}

type YoutubeWrapper struct {
	*youtube.Service
}

func (service YoutubeWrapper) VideoList(channelId string) map[string]video {
	flag.Parse()

	call := service.Activities.List("id,snippet").
		ChannelId(channelId).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call with id %v stack: %v", channelId, err)
	}

	videos := make(map[string]video)

	for _, item := range response.Items {
		videos[item.Id] = video{
			title:        item.Snippet.Title,
			description:  item.Snippet.Description,
			thumbnailURL: item.Snippet.Thumbnails.High.Url,
			channelTitle: item.Snippet.ChannelTitle,
			channelID:    item.Snippet.ChannelId}
	}

	return videos
}

func PrintIDs(sectionName string, matches map[string]video) {
	fmt.Printf("%v:\n", sectionName)
	for id, video := range matches {
		fmt.Printf("[%v] %v %v\n", id, video.title, video.thumbnailURL)
	}
	fmt.Printf("\n\n")
}
