package web

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed all:dist
var rawEmbed embed.FS

func Handler(public string) http.HandlerFunc {
	staticFS, _ := fs.Sub(rawEmbed, "dist")
	index, err := fs.ReadFile(staticFS, "index.html")
	if err != nil {
		panic(err)
	}
	embeddedFs := http.StripPrefix(public, http.FileServer(http.FS(staticFS)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.ContainsRune(r.URL.Path, '.') {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		} else {
			embeddedFs.ServeHTTP(w, r)
		}
	})
}
