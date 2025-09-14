package main

import (
	"diary/internal/middleware"
	"log"
	"net/http"
)

const addr = "localhost:8000"

func notFound (w *middleware.ResponseWriter, r *middleware.Request) error {
	http.NotFound(w, r.Http)
	return nil
}
var notFoundHandler = middleware.Use(notFound)

func createServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", notFoundHandler)
	serveMux.HandleFunc("GET /records", getRecords)
	serveMux.HandleFunc("POST /records", addRecord)
	serveMux.HandleFunc("DELETE /records/{id}", deleteRecord)
	return serveMux
}

func main() {
	serveMux := createServeMux()
	log.Printf("Listening to http://%s...\n", addr)
	err := http.ListenAndServe(addr, serveMux)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
}
