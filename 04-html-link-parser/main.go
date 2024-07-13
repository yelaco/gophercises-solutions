package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yelaco/link/parser"
)

func main() {
	s, err := os.ReadFile("ex1.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parser.ParseHTML(string(s)))
}
