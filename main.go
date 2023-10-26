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

	addBookHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		htmlStr := fmt.Sprintf("<div class='font-bold border p-2 mb-2 border-gray-400 shadow-md flex hover:bg-white hover:cursor-pointer'> <div class='mr-auto'> <p>%s</p> <p class='text-xs text-gray-400'>%s</p> </div> </div>", title, author)

		template, _ := template.New("newBook").Parse(htmlStr)

		template.Execute(w, nil)
	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-book/", addBookHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}