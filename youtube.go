package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/youtube/activities"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var channelID = flag.String("channel-id", "UCLRCoLpDImkbFw8D3qXFpUQ", "Channel Id exemple")

type Configuration struct {
	developerKey string
}

func main() {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("key")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	s := &activities.MyService{service}

	videos := s.VideoList(*channelID)
	activities.PrintIDs("Videos", videos)
}
