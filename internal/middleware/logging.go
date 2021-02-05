package middleware

import (
	"net/http"

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
		wrapped := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		log.Infof("HTTP request method=%s host=%s URI=%s remote=%s status=%d", r.Method, r.Host, r.RequestURI, r.RemoteAddr, wrapped.status)
	})
}
