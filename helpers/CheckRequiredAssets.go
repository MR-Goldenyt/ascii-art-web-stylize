package helpers

import (
	"fmt"
	"os"
)

func CheckRequiredAssets() error {
	assets := []string{
		"assets/templates/index.html",
		"assets/templates/404.html",
		"assets/templates/500.html",
		"assets/static/styles.css",
		"assets/static/dropdown.svg",
		"assets/static/matrix.mp4",
	}

	for _, asset := range assets {
		if _, err := os.Stat(asset); err != nil {
			return fmt.Errorf("missing or unreadable asset: %s", asset)
		}
	}
	return nil
}
