package main

import (
	"github.com/tannergabriel/learning-go/beginner-programs/HttpServer/server"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", server.NewServer()))
}
