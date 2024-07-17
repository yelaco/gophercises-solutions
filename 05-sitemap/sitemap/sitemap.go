package sitemap

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/yelaco/sitemap/parser"
)

type Url struct {
	Loc string `xml:"loc"`
}

type UrlSet struct {
	URLs []Url `xml:"urlset"`
}

func BuildSitemap(url string) string {
	urlQueue := []string{url}
	siteMap := map[string]struct{}{}

	// Using BFS to crawl the sites
	for len(urlQueue) > 0 {
		links, _ := crawlLinks(urlQueue[0])

		for _, link := range links {
			internalUrl, same := sameDomain(url, link.Href)
			if !same {
				continue
			}
			if _, ok := siteMap[internalUrl]; ok {
				continue
			}
			siteMap[internalUrl] = struct{}{}
			urlQueue = append(urlQueue, strings.TrimRight(internalUrl, "/"))
			fmt.Println(internalUrl)
		}

		urlQueue = urlQueue[1:]
	}

	return sitemapToXML(siteMap)
}

func sameDomain(url, href string) (string, bool) {
	if href == "" {
		return "", true
	}
	if strings.Contains(href, url) {
		return href, true
	} else if href[0] == '/' {
		return url + href, true
	} else {
		return "", false
	}
}

func crawlLinks(url string) ([]parser.Link, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []parser.Link{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []parser.Link{}, errors.New("request not ok")
	}

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return []parser.Link{}, err
	}

	links, err := parser.ParseHTML(string(bodyByte))
	if err != nil {
		return []parser.Link{}, err
	}

	return links, nil
}

func sitemapToXML(sm map[string]struct{}) string {
	return ""
}
