package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"encoding/json"
)

var getRecords = middleware.Use(
	func(w *middleware.ResponseWriter, r *middleware.Request) error {
		records, err := database.GetRecordsByUser(r.User)
		if err != nil {
			return err
		}
		recordsJson, _ := json.Marshal(records)
		w.Header().Set("Content-Type", "application/json")
		w.Write(recordsJson)
		return nil
	},
	corsMiddleware,
	middleware.BasicAuth,
)
