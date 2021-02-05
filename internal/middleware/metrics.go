package middleware

import (
	"net/http"
	"path"
	"time"

	"github.com/335is/server/internal/metrics"
)

// MetricsMiddleware implements mux.MiddlewareFunc.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		value := time.Since(start)

		label := r.Method + ":" + path.Join(r.RemoteAddr, r.RequestURI)
		metrics.AddInteger(label, 1)
		metrics.AddDuration(label, value)
	})
}
