package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	Title string
	Author string
}

func main() {
	fmt.Println("Hello, World!")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		books := map[string][]Book{
			"Books": {
				{Title: "Dune", Author: "Frank Herbert"},
				{Title: "Neuromancer", Author: "William Gibson"},
				{Title: "Foundation", Author: "Isaac Asimov"},
				{Title: "1984", Author: "George Orwell"},
				{Title: "The Expanse", Author: "James S.A. Corey"},
			},
		}
		templ.Execute(w, books)
	}

	http.HandleFunc("/", handler1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}