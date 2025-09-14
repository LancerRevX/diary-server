package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type ResponseWriter struct {
	http       http.ResponseWriter
	request    *Request
	statusCode int
	logMessage string
}

func (w *ResponseWriter) AddLog(format string, a ...any) {
	if w.logMessage != "" {
		w.logMessage += ": "
	}
	w.logMessage += fmt.Sprintf(format, a...)
}

func (w *ResponseWriter) Write(content []byte) (int, error) {
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

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.http.WriteHeader(statusCode)
}

func (w *ResponseWriter) Header() http.Header {
	return w.http.Header()
}
