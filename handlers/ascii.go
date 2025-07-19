package handlers

import (
	"ascii-art-web-stylize/helpers"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func PostArtHandler(tmpl *template.Template, fonts embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := helpers.CheckRequiredAssets(); err != nil {
			helpers.Handle500(w, tmpl, err)
			return
		}

		input := r.FormValue("text")
		banner := strings.ToLower(r.FormValue("banner"))
		showImage, imagePath := helpers.LuckyWinner()

		data := struct {
			Input     string
			Banner    string
			Result    string
			Linecount int
			Error     string

			// 1% chance of an image showing up
			ShowImage bool
			ImagePath string
		}{
			Input:  input,
			Banner: banner,
			ShowImage: showImage,
			ImagePath: imagePath,
		}

		// 1) Validate input
		if text := strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", ""), "\n", ""); strings.TrimSpace(input) == "" || len(text) > 300 {
			data.Error = "Invalid input."
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "index.html", data)
			return
		}

		// 1) Validate banner choice
		switch banner {
		case "standard", "shadow", "thinkertoy":
		default:
			data.Error = "Invalid banner font."
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "index.html", data)
			return
		}

		// 2) Load the font file
		fontPath := fmt.Sprintf("fonts/%s.txt", banner)
		raw, err := fonts.ReadFile(fontPath)
		if err != nil {
			data.Error = "Could not load font file."
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "index.html", data)
			return
		}
		// 3) replace all carriage return with newlines (with out this text is broken on windows computers)
		cleanFont := strings.ReplaceAll(string(raw), "\r\n", "\n")
		cleanFont = strings.ReplaceAll(cleanFont, "\r", "\n")
		fontLines := strings.Split(cleanFont, "\n")

		input = strings.ReplaceAll(input, "\r\n", "\n")
		input = strings.ReplaceAll(input, "\r", "\n")

		// 4) Split input into segments (handles real newlines and "\n")
		segments := helpers.SplitSegments(input)

		// 5) Validate characters exist in font
		rowsPerChar := 9
		numChars := len(fontLines) / rowsPerChar
		for _, seg := range segments {
			for _, rnr := range seg {
				idx := int(rnr) - 32
				if idx < 0 || idx >= numChars {
					data.Error = fmt.Sprintf("Invalid character %q in input.", rnr)
					w.WriteHeader(http.StatusBadRequest)
					tmpl.ExecuteTemplate(w, "index.html", data)
					return
				}
			}
		}

		data.Result, data.Linecount = helpers.RenderAscii(segments, fontLines)

		w.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(w, "index.html", data)
	}
}
