package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Joao", "joao.pedro@gmail.com"}

		if err := templates.ExecuteTemplate(w, "home.html", u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("template error:", err)
		}
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
