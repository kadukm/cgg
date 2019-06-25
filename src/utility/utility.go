package utility

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	TempFileName = "temp.png"
)

type IntTuple struct {
	A, B int
}

//Fill fills img with color c
func Fill(img draw.Image, c color.Color) {
	for xx := img.Bounds().Min.X; xx <= img.Bounds().Max.X; xx++ {
		for yy := img.Bounds().Min.Y; yy < img.Bounds().Max.Y; yy++ {
			img.Set(xx, yy, c)
		}
	}
}

//DrawHorizontalLine draws line from (xxStart, yy0) to (xxEnd, yy0)
func DrawHorizontalLine(img draw.Image, yy0, xxStart, xxEnd int, c color.Color) {
	for xx := xxStart; xx <= xxEnd; xx++ {
		img.Set(xx, yy0, c)
	}
}

//DrawVerticalLine draws line from (xx0, yyStart) to (xx0, yyEnd)
func DrawVerticalLine(img draw.Image, xx0, yyStart, yyEnd int, c color.Color) {
	for yy := yyStart; yy < yyEnd; yy++ {
		img.Set(xx0, yy, c)
	}
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
