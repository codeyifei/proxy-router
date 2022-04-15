package middleware

import (
	"log"
	"net/http"
)

func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("scheme: %s, host: %s, path: %s", r.URL.Scheme, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
