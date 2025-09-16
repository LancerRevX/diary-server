package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"encoding/json"
	"net/http"
	"time"
)

type addRecordJsonReq struct {
	Content string
}

type addRecordJsonRes struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

var addRecord = middleware.Use(
	func(req *middleware.MyRequest, res *middleware.MyResponse) error {
		

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

		responseObject := addRecordJsonRes{
			Id:        record.Id,
			CreatedAt: record.CreatedAt,
		}
		responseJson, _ := json.Marshal(responseObject)

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJson)

		return nil
	},
	middleware.RequireContentTypeJson,
	middleware.BasicAuth,
)
