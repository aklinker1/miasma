package middleware

import (
	"net/http"
	"strings"
)

func UI() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/api") {
				http.FileServer(http.Dir("./dashboard")).ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
