package middleware

import (
	"net/http"
	"strings"
)

const apiKeyName = "APIKEY"
const apiKeyGuest = "SERVER-APIKEY-473b29ba-4ab3-46fa-bda1-9015444d70b5"

// AuthMiddleware is middleware that handles auth verification on each request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

// Check for proper API key in request header or query parameters
// Could make a call to some external auth service (Vault, cred mgr, etc.)
func isAuthorized(r *http.Request) bool {

	// first look at request header
	k := r.Header.Get(apiKeyName)
	if strings.EqualFold(k, apiKeyGuest) {
		return true
	}

	// next check query parameter
	k = r.URL.Query().Get(apiKeyName)
	return strings.EqualFold(k, apiKeyGuest)
}
