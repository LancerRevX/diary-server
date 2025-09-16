package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"log"
	"net/http"
)

const addr = "localhost:8000"

func notFound(w *middleware.MyResponseWriter, r *middleware.MyRequest) error {
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

var corsMiddleware = middleware.AccessControlAllowOrigin("http://localhost:3000")

func main() {
	err := database.Open()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	serveMux := createServeMux()
	log.Printf("Listening to http://%s...\n", addr)
	err = http.ListenAndServe(addr, serveMux)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
}
