package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/youtube/activities"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type Configuration struct {
	developerKey string
}

type Channel struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func main() {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("key")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	s := &activities.YoutubeWrapper{service}
	ch := channels()
	for _, p := range ch {
		videos := s.VideoList(p.ID)
		activities.PrintIDs("Videos", videos)
	}
}

func channels() []Channel {
	raw, err := ioutil.ReadFile("./channels.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Channel
	json.Unmarshal(raw, &c)
	return c
}
