package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("story.gohtml"))
}

func main() {
	var storyArcMap map[string]StoryArc

	storyFile, err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(storyFile)
	if err := decoder.Decode(&storyArcMap); err != nil {
		panic(err)
	}

	advantureHandler := AdvantureHandler{
		StoryArcMap: storyArcMap,
	}

	http.Handle("GET /adventure", &advantureHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
