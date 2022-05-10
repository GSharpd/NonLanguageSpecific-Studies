package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Username string
	Email    string
}

func home(writer http.ResponseWriter, request *http.Request) {
	user := user{
		"Mad",
		"mad@gmail.com",
	}
	templates.ExecuteTemplate(writer, "home.html", user)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html")) // load all templates into templates variable

	http.HandleFunc("/home", home)

	fmt.Println("listening on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
