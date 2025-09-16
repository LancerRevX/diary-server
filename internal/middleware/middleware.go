package middleware

import (
	"diary/internal/database"
	"fmt"
	"net/http"
)

type MyHandlerFunc = func(req *MyRequest, res *MyResponse) error

type Middleware struct {
	Name string
	Func func(req *MyRequest, res *MyResponse) (bool, error)
}

func Use(
	handler MyHandlerFunc,
	middlewares ...Middleware,
) http.HandlerFunc {
	return func(httpW http.ResponseWriter, httpR *http.Request) {
		r := &MyRequest{Http: httpR}
		w := &MyResponseWriter{http: httpW, request: r, statusCode: http.StatusOK}
		for _, middleware := range middlewares {
			continue_, err := middleware.Func(w, r)
			if err != nil {
				w.AddLog("error executing middleware \"%s\": %s", middleware.Name, err)
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
