package middleware

import (
	"diary/internal/database"
	"encoding/json"
	"net/http"
)

type MyRequest struct {
	Http *http.Request
	User *database.User
}

func (req *MyRequest) Json(v any) (ok bool, err error) {
	decoder := json.NewDecoder(req.Http.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(v)
	if err != nil {
		res.W.AddLog("json decode error")
		http.Error(
			w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		)
		return nil
	}
}
