package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type MyResponseWriter struct {
	http       http.ResponseWriter
	request    *MyRequest
	statusCode int
	logMessage string
}

func (w *MyResponseWriter) AddLog(format string, a ...any) {
	if w.logMessage != "" {
		w.logMessage += ": "
	}
	w.logMessage += fmt.Sprintf(format, a...)
}

func (w *MyResponseWriter) Write(content []byte) (int, error) {
	result, err := w.http.Write(content)

	responseText := string(content)
	if len(responseText) > 30 {
		responseText = responseText[:30] + "..."
	}

	logMessage := fmt.Sprintf(
		"%s %s %s %s HTTP %d",
		w.request.Http.RemoteAddr,
		w.request.Http.UserAgent(),
		w.request.Http.Method,
		w.request.Http.URL,
		w.statusCode,
	)
	if w.request.User != nil {
		logMessage += fmt.Sprintf(" user \"%s\"", w.request.User.Login)
	}
	if w.logMessage != "" {
		logMessage += fmt.Sprintf(": %s", w.logMessage)
	}
	logMessage += fmt.Sprintf(" => %s", responseText)
	log.Print(logMessage)

	return result, err
}

func (w *MyResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.http.WriteHeader(statusCode)
}

func (w *MyResponseWriter) Header() http.Header {
	return w.http.Header()
}
