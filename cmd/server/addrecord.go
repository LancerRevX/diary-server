package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"encoding/json"
	"net/http"
	"time"
)

type recordJsonIn struct {
	Content string
}

type recordJsonOut struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

var addRecord = middleware.Use(
	func(w *middleware.ResponseWriter, r *middleware.Request) error {
		decoder := json.NewDecoder(r.Http.Body)
		decoder.DisallowUnknownFields()
		recordJson := recordJsonIn{}
		err := decoder.Decode(&recordJson)
		if err != nil {
			w.AddLog("json decode error")
			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)
			return nil
		}

		if recordJson.Content == "" {
			w.AddLog("empty content")
			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)
			return nil
		}

		record, err := database.CreateRecord(r.User, recordJson.Content)
		if err != nil {
			return err
		}

		responseObject := recordJsonOut{
			Id:        record.Id,
			CreatedAt: record.CreatedAt,
		}
		responseJson, _ := json.Marshal(responseObject)

		w.Write(responseJson)

		return nil
	},
	middleware.BasicAuth,
)
