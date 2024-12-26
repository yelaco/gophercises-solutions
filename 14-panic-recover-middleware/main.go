package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEV_MODE  = "development"
	PROD_MODE = "production"
)

var mode = os.Getenv("MODE")

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/panic/", panicDemo)
	router.HandleFunc("/panic-after/", panicAfterDemo)
	router.HandleFunc("/", hello)

	stack := createStack(panicRecover)

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(router),
	}
	log.Fatal(server.ListenAndServe())
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
