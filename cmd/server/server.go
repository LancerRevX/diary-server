package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"encoding/json"
	"log"
	"net/http"
)

const addr = "localhost:8000"

var notFound = middleware.With(
	func(w *middleware.ResponseWriter, r *middleware.Request) error {
		http.NotFound(w, r.Http)
		return nil
	},
)

var getRecords = middleware.With(
	func(w *middleware.ResponseWriter, r *middleware.Request) error {
		records, err := database.GetRecordsForUser(r.User)
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
	http.HandleFunc("/", notFound)
	http.HandleFunc("GET /records", getRecords)
	log.Printf("Listening to http://%s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
}
