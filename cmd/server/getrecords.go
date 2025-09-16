package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
)

var getRecords = middleware.Use(
	func(req *middleware.MyRequest, res *middleware.MyResponse) error {
		records, err := database.GetRecordsByUser(req.User.Id)
		if err != nil {
			return err
		}
		err = res.Json(records)
		if err != nil {
			return err
		}
		return nil
	},
	corsMiddleware,
	middleware.BasicAuth,
)
