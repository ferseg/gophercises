package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ferseg/gophercises/cyoa/src"
)

func main() {
	filename := flag.String("filename", "stories.json", "The name of the file that contains the stories")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}
	defer file.Close()

  story, err := src.ReadJsonStory(file)
  if err != nil {
    log.Fatal("Error reading file content: ", err)
	}
  h := src.NewHandler(story)
  http.ListenAndServe(":8080", h)
}
