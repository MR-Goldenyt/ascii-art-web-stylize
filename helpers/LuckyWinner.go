package helpers

import (
	"ascii-art-web-stylize/assets"
	"io/fs"
	"math/rand"
	"strings"
	"time"
)

func LuckyWinner() (showImage bool, ImagePath string) {
	images := getImageList()

	
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 1% chance to show
	if len(images) > 0 && rng.Float64() < 0.01 {
		return true, images[rng.Intn(len(images))]
	}
	
	return false, ""
}

func getImageList() []string {
	var images []string
	fs.WalkDir(assets.CssFS, "static/images", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(path), ".png") {
			images = append(images, "/"+path) // make it usable in <img src>
		}
		return nil
	})
	return images
}
