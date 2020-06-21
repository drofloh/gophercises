package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drofloh/gophercises/cyoa"
)

type Story map[string]Section

type Section struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	port := flag.Int("p", 3030, "The port to listen for requests on")
	filename := flag.String("f", "gopher.json", "The file json to load which contains the stories")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
