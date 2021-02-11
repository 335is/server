package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/335is/log"
)

// PanicMiddleware handles a panic, logs it and the stack trace, and recovers.
func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Errorf("%v", err)
				log.Errorf(string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
