package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHTML(s string) ([]Link, error) {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return nil, err
	}
	links := extractLinks(doc)

	return links, nil
}

func extractLinks(n *html.Node) []Link {
	links := []Link{}

	if n.Type == html.ElementNode && n.Data == "a" {
		link := Link{}
		for _, a := range n.Attr {
			if a.Key == "href" {
				link.Href = strings.TrimSpace(a.Val)
				break
			}
		}

		link.Text = extractText(n)

		links = append(links, link)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c)...)
	}

	return links
}

func extractText(n *html.Node) string {
	return ""
}
