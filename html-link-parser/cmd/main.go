package main

import (
	"flag"
	"fmt"
	"log"
	"os"
  "github.com/ferseg/gophercises/html-link-parser"
)


func main() {
  fileName := flag.String("file", "example.html", "The name of the file")
  flag.Parse()

  file, err := os.Open(*fileName)
  if err != nil {
    log.Fatal("Could not open file", err)
  }
  links, err := parser.Parse(file)
  if err != nil {
    log.Fatal("Could not parse file", err)
  }
  fmt.Print(links)
}

