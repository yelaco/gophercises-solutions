package main

import (
	"flag"
	"fmt"

	"github.com/yelaco/sitemap/sitemap"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "https://monstar-lab.com/vn/", "URL to build a sitemap from")
	flag.Parse()

	fmt.Println(sitemap.BuildSitemap(url))
}
