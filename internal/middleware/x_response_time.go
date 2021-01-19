package middleware

import (
	"net/http"
	"time"
)

func XResponseTime() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			w.Header().Add("Trailer", "X-Response-Time")
			next.ServeHTTP(w, r)
			elapsed := time.Since(start)
			w.Header().Add("X-Response-Time", elapsed.String())
		})
	}
}
