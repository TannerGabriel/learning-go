package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", HelloHandler)
	r.HandleFunc("/", HelloWorldHandler)
	return r
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name, exists := mux.Vars(r)["name"]

	if !exists {
		name = "world"
	}

	w.Write([]byte("Hello, " + name))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
