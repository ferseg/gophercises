package parser

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"golang.org/x/net/html"
)

var stringNode = `
<div>
  <div>
    <a href="/first-child">I'm the first</a>
  </div>
  <a href="/next-sibling>I'm the sibling</a>
</div>`

var mainNode = html.Node{
	FirstChild: &html.Node{
		Type:        html.ElementNode,
		Data:        "div",
		FirstChild:  createLinkElement("/first-child", "I'm the first"),
		NextSibling: createLinkElement("/next-sibling", "I'm the sibling"),
	},
}

func createLinkElement(href, text string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: href,
			},
		},
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: text,
		},
	}
}

func TestParse(t *testing.T) {
	r := strings.NewReader(stringNode)
	links, err := Parse(r)
	assert.Equal(t, nil, err)
	validateLinks(t, links)
}

func TestExtractLinkNodes(t *testing.T) {
	resultNodes := extractLinkNodes(&mainNode)
	assert.Equal(t, len(resultNodes), 2)
	for _, node := range resultNodes {
		assert.Equal(t, node.Data, "a")
	}
}

func validateLinks(t *testing.T, links []Link) {
	linkMap := map[string]string{
		"/first-child":  "I'm the first",
		"/next-sibling": "I'm the sibling",
	}

	for _, link := range links {
		text := linkMap[link.Href]
		assert.Equal(t, link.Text, text)
	}
}

func TestConvertToLink(t *testing.T) {
	nodes := extractLinkNodes(&mainNode)
	links := convertToLinks(nodes)
	assert.Equal(t, len(links), len(nodes))
	validateLinks(t, links)
}
