package main

import (
	"log"
	"net/http"
)

func user(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Loading user page..."))
}

func main() {
	// Communication protocol, it is the foundation of web communication
	// Client requests <-> Server responds

	// Routes
	// URI - Identifies the resource
	// Method - Http Verbs - What is done with said resource

	// Creating a server in Go

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is the response"))
	})

	http.HandleFunc("/user", user)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
