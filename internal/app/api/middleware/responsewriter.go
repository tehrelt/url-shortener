package middleware

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	code int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
