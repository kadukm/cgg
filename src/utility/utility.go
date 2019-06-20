package utility

import (
	"image/color"
	"image/draw"
)

func Fill(img draw.Image, c color.Color) {
	for xx := img.Bounds().Min.X; xx <= img.Bounds().Max.X; xx++ {
		for yy := img.Bounds().Min.Y; yy < img.Bounds().Max.Y; yy++ {
			img.Set(xx, yy, c)
		}
	}
}
