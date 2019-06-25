package utility

import (
	"image"
	"image/png"
	"log"
	"os"
)

const (
	TempFileName = "temp.png"
)

type Point struct {
	XX, YY int
}

func PointInsideImage(img image.Image, xx, yy int) bool {
	return (img.Bounds().Min.X <= xx && xx <= img.Bounds().Max.X &&
		img.Bounds().Min.Y <= yy && yy <= img.Bounds().Max.Y)
}

func SavePNG(img image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
