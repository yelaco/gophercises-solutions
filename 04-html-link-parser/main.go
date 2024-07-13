package main

import (
	"fmt"

	"github.com/yelaco/link/parser"
)

func main() {
	// ex1, err := os.ReadFile("ex1.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ex2, err := os.ReadFile("ex2.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ex3, err := os.ReadFile("ex3.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ex4, err := os.ReadFile("ex4.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ex5 := `<a href="#">Something here <a href="/dog">nested dog link</a></a>`

	// fmt.Println(parser.ParseHTML(string(ex1)))
	// fmt.Println(parser.ParseHTML(string(ex2)))
	// fmt.Println(parser.ParseHTML(string(ex3)))
	// fmt.Println(parser.ParseHTML(string(ex4)))
	fmt.Println(parser.ParseHTML(string(ex5)))
}
