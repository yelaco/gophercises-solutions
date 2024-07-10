package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/yelaco/url-shortener/urlshort"
)

var yamlFlag = flag.String("yaml", "", "specify yaml file for mapping paths to urls")
var jsonFlag = flag.String("json", "", "specify json file for mapping paths to urls")

func main() {
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var betterHandler http.Handler

	if *yamlFlag != "" {
		yaml, err := os.ReadFile(*yamlFlag)
		if err != nil {
			fmt.Println("Coulnd't read yaml file")
			os.Exit(1)
		}

		// Build the YAMLHandler using the mapHandler as the fallback
		betterHandler, err = urlshort.YAMLHandler(yaml, mapHandler)
		if err != nil {
			fmt.Println("Coulnd't parse yaml")
			os.Exit(1)
		}
	} else if *jsonFlag != "" {
		json, err := os.ReadFile(*jsonFlag)
		if err != nil {
			fmt.Println("Coulnd't read json file")
			os.Exit(1)
		}

		betterHandler, err = urlshort.JSONHandler(json, mapHandler)
		if err != nil {
			fmt.Println("Coulnd't parse json")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: -yaml or -json is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", betterHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
