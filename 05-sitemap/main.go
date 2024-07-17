package main

import (
	"flag"

	"github.com/yelaco/sitemap/sitemap"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "https://monstar-lab.com/vn/", "URL to build a sitemap from")
	flag.Parse()

	sitemap.BuildSitemap(url)
}
