package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs all requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Process the request
		next.ServeHTTP(w, r)

		// Log after processing
		duration := time.Since(start)
		log.Printf("[%s] %s %s %s", r.Method, r.URL.Path, r.Proto, duration)
	})
}

// ResponseWriter wraps http.ResponseWriter to capture status code
type ResponseWriter struct {
	http.ResponseWriter
	Status int
	Body   []byte
}

// NewResponseWriter creates a new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}

// WriteHeader captures the status code
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.Status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the response body
func (rw *ResponseWriter) Write(data []byte) (int, error) {
	rw.Body = data
	return rw.ResponseWriter.Write(data)
}
