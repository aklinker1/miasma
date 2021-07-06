package middleware

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/aklinker1/miasma/internal/shared/log"
)

//go:embed web/*
var embeddedFS embed.FS

func UI() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		webDir, err := fs.Sub(embeddedFS, "web")
		if err != nil {
			log.W("Failed to start dashboard:", err.Error())
		}
		webDirServer := http.StripPrefix("/", http.FileServer(http.FS(webDir)))

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			index, err := fs.ReadFile(webDir, "index.html")
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte{})
				return
			}
			if strings.HasPrefix(r.URL.Path, "/api") {
				next.ServeHTTP(w, r)
				return
			}
			if strings.ContainsRune(r.URL.Path, '.') {
				webDirServer.ServeHTTP(w, r)
				return
			}

			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
			// if !strings.HasPrefix(r.URL.Path, "/api") {
			// 	http.FileServer(http.Dir("./dashboard")).ServeHTTP(w, r)
			// 	return
			// }

			// next.ServeHTTP(w, r)
		})
	}
}
