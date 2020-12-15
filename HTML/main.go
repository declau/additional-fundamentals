package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Dunha",
		"dunha@gmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
