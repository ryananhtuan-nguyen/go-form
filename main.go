package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := []Film{
			{Title: "Inception", Director: "Christopher Nolan"},
			{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
			{Title: "Spirited Away", Director: "Hayao Miyazaki"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/add-film/", h2)

	fmt.Println("Server started and listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
