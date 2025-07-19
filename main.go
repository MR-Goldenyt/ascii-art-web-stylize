package main

import (
	"ascii-art-web-stylize/assets"
	"ascii-art-web-stylize/handlers"
	"ascii-art-web-stylize/helpers"
	"html/template"
	"log"
	"net/http"
)

func main() {
	err := helpers.CheckRequiredAssets()
	if err != nil {
		log.Printf("Necessary files are missing")
	}
	// Serve static files behind a custom handler
	http.Handle("/static/", handlers.AssetHandler())

	tmpl := template.Must(template.ParseFS(assets.TmplFS, "templates/*.html"))

	http.HandleFunc("/", handlers.GetHomeHandler(tmpl))
	http.HandleFunc("/ascii-art", handlers.PostArtHandler(tmpl, assets.FontsFS))

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
