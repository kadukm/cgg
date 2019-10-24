package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/kadukm/cgg/src/utility"
)

const (
	a, b float64 = -3, 3
	c, d float64 = -3, 3

	width  = 1000
	height = 600

	xLinesCount = 50
	stepsCount  = width * 2

	fileName = "fourthtask.png"
)

func f(x, y float64) float64 {
	return math.Cos(x * y)
}

func main() {
	fg := initFunctionGraph()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	drawFunction(img, fg)

	utility.SavePNG(img, fileName)
}

func initFunctionGraph() (fg utility.FunctionGraph3d) {
	fg.XMin, fg.XMax = a, b
	fg.YMin, fg.YMax = c, d
	fg.Width = width
	fg.Height = height
	fg.LinesCount = xLinesCount
	fg.StepsCount = stepsCount

	findMinAndMaxProjectionCoordinates(&fg)

	return
}

func findMinAndMaxProjectionCoordinates(fg *utility.FunctionGraph3d) {
	xxProjectionMin := math.MaxFloat64
	xxProjectionMax := -math.MaxFloat64
	yyProjectionMin := math.MaxFloat64
	yyProjectionMax := -math.MaxFloat64
	for i := 0; i <= fg.LinesCount; i++ {
		x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.LinesCount)
		for j := 0; j <= fg.StepsCount; j++ {
			y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.StepsCount)
			z := f(x, y)

			xxProjection := fg.GetXXProjection(x, y, z)
			yyProjection := fg.GetYYProjection(x, y, z)
			if xxProjection < xxProjectionMin {
				xxProjectionMin = xxProjection
			}
			if xxProjection > xxProjectionMax {
				xxProjectionMax = xxProjection
			}
			if yyProjection < yyProjectionMin {
				yyProjectionMin = yyProjection
			}
			if yyProjection > yyProjectionMax {
				yyProjectionMax = yyProjection
			}
		}
	}

	fg.XXProjectionMin, fg.XXProjectionMax = xxProjectionMin, xxProjectionMax
	fg.YYProjectionMin, fg.YYProjectionMax = yyProjectionMin, yyProjectionMax
}

func drawFunction(img draw.Image, fg utility.FunctionGraph3d) {
	bottomColor := color.RGBA{R: 0, G: 200, B: 255, A: 255}
	topColor := color.RGBA{R: 255, G: 0, B: 155, A: 255}
	bottom, top := initBottomAndTopArrays(fg)

	for i := 0; i <= fg.LinesCount; i++ {
		//TODO: remove code repeat
		x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.LinesCount)
		for j := 0; j <= fg.StepsCount; j++ {
			y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.StepsCount)
			z := f(x, y)

			xxProjection := fg.GetXXProjection(x, y, z)
			yyProjection := fg.GetYYProjection(x, y, z)

			xx := fg.XXProjectionToScreen(xxProjection)
			yy := fg.YYProjectionToScreen(yyProjection)
			if yy > bottom[xx] {
				bottom[xx] = yy
				img.Set(xx, yy, bottomColor)
			}
			if yy < top[xx] {
				top[xx] = yy
				img.Set(xx, yy, topColor)
			}
		}
	}

	_, top = initBottomAndTopArrays(fg)
	for j := 0; j <= fg.LinesCount; j++ {
		y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.LinesCount)
		for i := 0; i <= fg.StepsCount; i++ {
			x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.StepsCount)
			z := f(x, y)

			xxProjection := fg.GetXXProjection(x, y, z)
			yyProjection := fg.GetYYProjection(x, y, z)

			xx := fg.XXProjectionToScreen(xxProjection)
			yy := fg.YYProjectionToScreen(yyProjection)
			if yy > bottom[xx] {
				bottom[xx] = yy
				img.Set(xx, yy, bottomColor)
			}
			if yy < top[xx] {
				top[xx] = yy
				img.Set(xx, yy, topColor)
			}
		}
	}
}

func initBottomAndTopArrays(fg utility.FunctionGraph3d) ([]int, []int) {
	top := make([]int, fg.Width+1)
	bottom := make([]int, fg.Width+1)
	for i := 0; i < fg.Width; i++ {
		top[i] = fg.Height + 1
		bottom[i] = -1
	}

	return bottom, top
}
