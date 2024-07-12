package main

import (
	"log"
	"net/http"
)

type AdvantureHandler struct {
	StoryArcMap map[string]StoryArc
}

func (av *AdvantureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Query().Get("arc")

	if storyArc, exists := av.StoryArcMap[arc]; exists {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		if err := tpl.ExecuteTemplate(w, "story.gohtml", storyArc); err != nil {
			log.Println(err.Error())
		}
	} else {
		respondWithError(w, http.StatusNotFound, "Arc not exists")
	}
}
