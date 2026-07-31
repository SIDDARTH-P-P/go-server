package auth

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
)

const (
	Username = "admin"
	Password = "password123"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			writeAuthError(w)
			return
		}

		if subtle.ConstantTimeCompare([]byte(user), []byte(Username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(Password)) != 1 {
			writeAuthError(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeAuthError(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error":   "unauthorized: invalid or missing basic auth credentials",
	})
}
