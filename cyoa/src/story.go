package src

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var DefaultHtmlTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Choose your own advendure</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
  </head>
  <body>
    <h1>{{.Title}}</h1>
    <div>
      {{range .Paragraphs}}
      <p>{{.}}</p>
      {{end}}
    </div>
    <div>
      <ul>
        {{range .Options}}
        <li><a href="/{{.Section}}">{{.Text}}</a></li>
        {{end}}
      </ul>
    </div>
  </body>
</html>`

type Story map[string]StorySection

type StoryOption struct {
	Text    string `json:"text"`
	Section string `json:"arc"`
}

type StorySection struct {
	Title      string        `json:"title"`
	Paragraphs []string      `json:"story"`
	Options    []StoryOption `json:"options"`
}

func ReadJsonStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type handler struct {
	story Story
}

func NewHandler(story Story) http.Handler {
	return handler{story}
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(DefaultHtmlTmpl))
  path := strings.TrimSpace(r.URL.Path)
  log.Printf("The path is: %s", path)
  if path == "" || path == "/" {
    path = "/intro"
  } 
  path = path[1:]
  err := tpl.Execute(rw, h.story[path])
	if err != nil {
		log.Fatal("Error parsing the template", err)
	}

}
