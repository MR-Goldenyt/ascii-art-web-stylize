package helpers

import (
	"ascii-art-web-stylize/assets"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Handle500(w http.ResponseWriter, tmpl *template.Template, err1 error) {
	if _, err := fs.Stat(assets.TmplFS, "templates/500.html"); err != nil {
		log.Printf("500 internal error: %v", err1)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("500 internal error: %v", err1)
	w.WriteHeader(http.StatusInternalServerError)
	tmpl.ExecuteTemplate(w, "500.html", http.StatusInternalServerError)
	return

}
