package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Associate URLs requested to functions that handle the requests
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/headers", headers)

	// Start the web server
	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Basic handler for /hello requests
func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	return
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
