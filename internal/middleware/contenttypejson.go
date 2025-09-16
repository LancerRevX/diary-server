package middleware

import "net/http"

var RequireContentTypeJson Middleware = Middleware{
	Name: "Json",
	Func: func(w *MyResponseWriter, r *MyRequest) (bool, error) {
		contentType := r.Http.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.AddLog("invalid content type: %s", contentType)
			http.Error(
				w,
				http.StatusText(http.StatusUnsupportedMediaType),
				http.StatusUnsupportedMediaType,
			)
			return false, nil
		}

		return true, nil
	},
}
