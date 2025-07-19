// main.go
package handlers

import (
	"ascii-art-web-stylize/helpers"
	"html/template"
	"net/http"
)

func GetHomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
			helpers.Handle404(w, tmpl)
			return
		}

		if err := helpers.CheckRequiredAssets(); err != nil {
			helpers.Handle500(w, tmpl, err)
			return
		}

		showImage, imagePath := helpers.LuckyWinner()

		data := struct {
			Input  string
			Banner string
			Result string
			Error  string

			// 1% chance of an image showing up
			ShowImage bool
			ImagePath string
		}{
			Banner: "standard",
			ShowImage: showImage,
			ImagePath: imagePath,

		}
		w.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(w, "index.html", data)
	}
}
