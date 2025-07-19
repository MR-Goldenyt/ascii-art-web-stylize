package assets

import "embed"

//go:embed templates/*.html
var TmplFS embed.FS

//go:embed static/*
var CssFS embed.FS

//go:embed fonts/*.txt
var FontsFS embed.FS

