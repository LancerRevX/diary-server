package middleware

import "fmt"

func AccessControlAllowOrigin(origin string) Middleware {
	return Middleware{
		Name: fmt.Sprintf(`Access-Control-Allow-Origin="%s"`, origin),
		Func: func(w *ResponseWriter, r *Request) (bool, error) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			return true, nil
		},
	}
}