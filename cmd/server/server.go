package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"encoding/json"
	"log"
	"net/http"
)

const addr = "localhost:8000"

func notFound (w *middleware.ResponseWriter, r *middleware.Request) error {
	http.NotFound(w, r.Http)
	return nil
}
var notFoundHandler = middleware.Use(notFound)

var getRecords = middleware.Use(
	func(w *middleware.ResponseWriter, r *middleware.Request) error {
		records, err := database.GetRecords(r.User)
		if err != nil {
			return err
		}
		recordsJson, err := json.Marshal(records)
		w.Write(recordsJson)
		return nil
	},
	middleware.BasicAuth,
)

func main() {
	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("GET /records", getRecords)
	http.HandleFunc("POST /records", addRecord)
	http.HandleFunc("DELETE /records/{id}", deleteRecord)
	log.Printf("Listening to http://%s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
}
