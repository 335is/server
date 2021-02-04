package router

import (
	"net/http"
	"time"

	"github.com/335is/log"
)

// responseWriter wraps http.ResponseWriter allowing the HTTP status code to be captured for logging.
// status needs to be initialized to 200 (OK), because WriteHeader() isn't called by http.ResponseWriter code.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader is only called from our routing code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)

	return
}

// LoggingMiddleware is middleware that logs each request, even ones that don't match a route
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		dur := time.Since(start)
		log.Infof("HTTP request method=%s host=%s URI=%s remote=%s status=%d duration=%s", r.Method, r.Host, r.RequestURI, r.RemoteAddr, wrapped.status, dur.String())
	})
}
