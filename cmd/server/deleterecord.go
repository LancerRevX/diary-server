package main

import (
	"diary/internal/database"
	"diary/internal/middleware"
	"net/http"
	"strconv"
)

var deleteRecord = middleware.Use(
	func(w *middleware.MyResponseWriter, r *middleware.MyRequest) error {
		recordId, err := strconv.ParseInt(r.Http.PathValue("id"), 10, 64)
		if err != nil {
			w.AddLog("invalid path parameter: %s", r.Http.PathValue("id"))
			return notFound(w, r)
		}

		authorized, err := r.User.HasRecord(recordId)
		if err != nil {
			return err
		}
		if !authorized {
			w.AddLog("tried to delete record with id \"%d\"", recordId)
			http.Error(
				w,
				http.StatusText(http.StatusForbidden),
				http.StatusForbidden,
			)
			return nil
		}

		err = database.DeleteRecord(recordId)
		if err != nil {
			return err
		}

		w.AddLog("deleted record with id \"%d\"", recordId)
		w.Write([]byte(http.StatusText(http.StatusOK)))

		return nil
	},
	middleware.BasicAuth,
)
