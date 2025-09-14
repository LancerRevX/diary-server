package middleware

import (
	"diary/internal/database"
	"net/http"
)

func sendUnauthorized(w http.ResponseWriter) {
	// this header entry is needed for a browser basic auth popup
	w.Header().Set("WWW-Authenticate", "Basic")
	
	http.Error(
		w,
		http.StatusText(http.StatusUnauthorized),
		http.StatusUnauthorized,
	)
}

var BasicAuth = Middleware{
	Name: "Basic Auth",
	Func: func(w *ResponseWriter, r *Request) (bool, error) {
		login, password, ok := r.Http.BasicAuth()
		if !ok {
			w.logMessage = "no basic auth credentials provided"
			sendUnauthorized(w)
			return false, nil
		}

		user, err := database.ValidateCredentials(login, password)
		switch err {
		case nil:
			break
		case database.ErrInvalidPassword:
			fallthrough
		case database.ErrUserNotFound:
			w.AddLog("invalid user \"%s\" with password \"%s\"",login,password)
			sendUnauthorized(w)
			return false, nil
		default:
			return false, err
		}

		r.User = user
		return true, nil
	},
}
