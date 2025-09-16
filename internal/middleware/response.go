package middleware

import "encoding/json"

type MyResponse struct {
	W *MyResponseWriter
}

func (res *MyResponse) Json(v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	res.W.Header().Set("Content-Type", "application/json")
	res.W.Write(body)
	return nil
}