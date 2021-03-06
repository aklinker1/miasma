package middleware

import (
	"net/http"
	"time"

	"github.com/aklinker1/miasma/internal/server/utils/env"
	"github.com/aklinker1/miasma/internal/shared/log"
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
			start := time.Now()
			recorder := &StatusRecorder{
				ResponseWriter: w,
				Status:         200,
			}
			log.D("<<<<<< %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(recorder, r)
			log.D(">>>>>> %d (%s)", recorder.Status, time.Since(start).String())
		})
	}
}
