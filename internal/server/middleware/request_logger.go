package middleware

import (
	"net/http"

	"github.com/aklinker1/miasma/internal/server/utils/env"
	"github.com/aklinker1/miasma/internal/server/utils/log"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func RequestLogger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if env.IS_PROD {
			return next
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &StatusRecorder{
				ResponseWriter: w,
				Status:         200,
			}
			log.V("<<<<<< %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(recorder, r)
			log.V(">>>>>> %d (%s)", recorder.Status, w.Header().Get("X-Response-Time"))
		})
	}
}
