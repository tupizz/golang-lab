package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Name  string
	Email string
}

type homeModelData struct {
	Usuario usuario
	Valid   bool
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			Name:  "Tadeu",
			Email: "tadeu.tupiz@gmail.com",
		}

		model := homeModelData{
			Usuario: u,
			Valid:   true,
		}

		templates.ExecuteTemplate(w, "home.html", model)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "about.html", nil)
	})

	log.Println("listening on port 5500")
	log.Fatal(http.ListenAndServe(":5500", nil))
}
