package main

import (
	"image"
	"image/color"
	"math"

	"github.com/kadukm/cgg/src/utility"
	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	width  = 1000
	height = 600

	a float64 = -100
	b float64 = 100

	fileName = "firsttask.png"
)

func f(x float64) float64 {
	return math.Sin(x) * x
}

func cartesianToScreen() (int, int) {
	return 0, 0
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	utility.Fill(img, color.White)

	gc := draw2dimg.NewGraphicContext(img)
	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(1)
	gc.BeginPath()

	yMin := f(a)
	yMax := yMin

	for xx := 0; xx <= width; xx++ {
		x := a + float64(xx)*(b-a)/width
		y := f(x)
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
	}

	yyFloat := (f(a) - yMax) * height / (yMin - yMax)
	gc.MoveTo(0, yyFloat)

	for xx := 1; xx <= width; xx++ {
		x := a + float64(xx)*(b-a)/width
		y := f(x)
		yyFloat = (y - yMax) * height / (yMin - yMax)
		yy := int(yyFloat)
		gc.LineTo(float64(xx), float64(yy))
	}

	gc.Stroke()

	draw2dimg.SaveToPngFile(fileName, img)
}
