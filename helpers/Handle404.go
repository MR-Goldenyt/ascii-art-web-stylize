package helpers

import (
	"ascii-art-web-stylize/assets"
	"html/template"
	"io/fs"
	"net/http"
)

func Handle404(w http.ResponseWriter, tmpl *template.Template) {
	if _, err := fs.Stat(assets.TmplFS, "templates/404.html"); err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	tmpl.ExecuteTemplate(w, "404.html", http.StatusNotFound)
	return

}
