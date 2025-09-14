package middleware

import (
	"diary/internal/database"
	"fmt"
	"net/http"
)

type Request struct {
	Http *http.Request
	User *database.User
}

type HandlerFunc = func(w *ResponseWriter, r *Request) error

type Middleware struct {
	Name string
	Func func(w *ResponseWriter, r *Request) (bool, error)
}

func With(
	handler HandlerFunc,
	middlewares ...Middleware,
) http.HandlerFunc {
	return func(httpW http.ResponseWriter, httpR *http.Request) {
		r := &Request{Http: httpR}
		w := &ResponseWriter{http: httpW, request: r, statusCode: http.StatusOK}
		for _, middleware := range middlewares {
			continue_, err := middleware.Func(w, r)
			if err != nil {
				w.AddLog(fmt.Sprintf(
					"error executing middleware \"%s\": %s",
					middleware.Name,
					err,
				))
				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
				return
			}
			if !continue_ {
				return
			}
		}
		err := handler(w, r)
		if err != nil {
			w.AddLog(fmt.Sprintf("error executing handler: %s", err))
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
	}
}
