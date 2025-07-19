package handlers

import (
	"ascii-art-web-stylize/assets"
	"net/http"
	"strings"
)

func AssetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/static/")
		data, err := assets.CssFS.ReadFile("static/" + path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		switch {
		case strings.HasSuffix(path, ".css"):
			w.Header().Set("Content-Type", "text/css")
		case strings.HasSuffix(path, ".mp4"):
			w.Header().Set("Content-Type", "video/mp4")
		}

		w.Write(data)
	})
}