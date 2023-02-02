package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
  Href string
  Text string
}


func Parse(r io.Reader) ([]Link, error) {
  doc, err := html.Parse(r)
  if err != nil {
    return nil, err
  }
  nodes := extractLinkNodes(doc)
  links := convertToLinks(nodes)
  return links, nil
}

func convertToLinks(nodes []*html.Node) []Link {
  var links []Link
  for _, node := range nodes {
    var link Link
    for _, attr := range node.Attr {
      if attr.Key == "href" {
        link.Href = attr.Val
        break;
      }
    }
    link.Text = extractTextFromNode(node)
    links = append(links, link)
  }
  return links
}

func extractTextFromNode(node *html.Node) string {
  if node.Type == html.TextNode {
    return node.Data
  }

  if node.Type != html.ElementNode {
    return ""
  }

  var result string
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    result += extractTextFromNode(c)
  }
  return strings.Join(strings.Fields(result), " ")
}

func extractLinkNodes(node *html.Node) []*html.Node {
  if node.Type == html.ElementNode && node.Data == "a" {
    return []*html.Node{node}
  }
	var ret []*html.Node
  for c := node.FirstChild; c != nil; c = c.NextSibling {
    ret = append(ret, extractLinkNodes(c)...)
  }
  return ret
}
