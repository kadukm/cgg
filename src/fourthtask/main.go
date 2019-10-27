package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/kadukm/cgg/src/utility"
)

const (
	x1, x2 float64 = -3, 3
	y1, y2 float64 = -3, 3

	width  = 1000
	height = 600

	linesCount       = 50
	stepsByLineCount = width * 2

	fileName = "fourthtask.png"
)

func f(x, y float64) float64 {
	return math.Cos(x * y)
	//return x * math.Pow(y, 3) - y * math.Pow(x, 3)
}

func main() {
	fg := initFunctionGraph()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	drawFunction(img, fg)

	utility.SavePNG(img, fileName)
}

func initFunctionGraph() (fg utility.FunctionGraph3D) {
	fg.XMin, fg.XMax = x1, x2
	fg.YMin, fg.YMax = y1, y2
	fg.Width = width - 1
	fg.Height = height - 1
	fg.LinesCount = linesCount
	fg.StepsByLineCount = stepsByLineCount

	findMinAndMaxProjectionCoordinates(&fg)

	return
}

func findMinAndMaxProjectionCoordinates(fg *utility.FunctionGraph3D) {
	xxProjectionMin := math.MaxFloat64
	xxProjectionMax := -math.MaxFloat64
	yyProjectionMin := math.MaxFloat64
	yyProjectionMax := -math.MaxFloat64
	for i := 0; i <= fg.LinesCount; i++ {
		x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.LinesCount)
		for j := 0; j <= fg.StepsByLineCount; j++ {
			y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.StepsByLineCount)
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

func drawFunction(img draw.Image, fg utility.FunctionGraph3D) {
	bottomColor := color.RGBA{R: 0, G: 200, B: 255, A: 255}
	topColor := color.RGBA{R: 255, G: 0, B: 155, A: 255}

	initTop := drawFirstTopLines(img, fg, topColor)
	drawFunctionMovingByX(img, fg, initTop, bottomColor, topColor)
	drawFunctionMovingByY(img, fg, initTop, bottomColor, topColor)
}

func drawFirstTopLines(img draw.Image, fg utility.FunctionGraph3D, topColor color.Color) []int {
	top := utility.InitIntSlice(fg.Width+1, fg.Height)
	topInitialized := make([]bool, fg.Width+1)
	drawFirstXY := func(x, y float64) {
		z := f(x, y)

		xxProjection := fg.GetXXProjection(x, y, z)
		yyProjection := fg.GetYYProjection(x, y, z)

		xx := fg.XXProjectionToScreen(xxProjection)
		yy := fg.YYProjectionToScreen(yyProjection)
		img.Set(xx, yy, topColor)
		if !topInitialized[xx] || yy > top[xx] {
			topInitialized[xx] = true
			top[xx] = yy
		}
	}
	x := fg.XMax
	for j := 0; j <= fg.StepsByLineCount; j++ {
		y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.StepsByLineCount)
		drawFirstXY(x, y)
	}

	y := fg.YMax
	for i := 0; i <= fg.StepsByLineCount; i++ {
		x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.StepsByLineCount)
		drawFirstXY(x, y)
	}

	return top
}

func drawFunctionMovingByX(img draw.Image, fg utility.FunctionGraph3D, initTop []int,
		bottomColor color.Color, topColor color.Color) {
	bottom := make([]int, fg.Width+1)
	top := make([]int, fg.Width+1)
	copy(bottom, initTop)
	copy(top, initTop)
	for i := 1; i <= fg.LinesCount; i++ {
		x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.LinesCount)
		for j := 0; j <= fg.StepsByLineCount; j++ {
			y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.StepsByLineCount)
			drawXY(img, fg, x, y, bottom, top, bottomColor, topColor)
		}
	}
}

func drawFunctionMovingByY(img draw.Image, fg utility.FunctionGraph3D, initTop []int,
		bottomColor color.Color, topColor color.Color) {
	bottom := make([]int, fg.Width+1)
	top := make([]int, fg.Width+1)
	copy(bottom, initTop)
	copy(top, initTop)
	for j := 1; j <= fg.LinesCount; j++ {
		y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.LinesCount)
		for i := 0; i <= fg.StepsByLineCount; i++ {
			x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.StepsByLineCount)
			drawXY(img, fg, x, y, bottom, top, bottomColor, topColor)
		}
	}
}

func drawXY(img draw.Image, fg utility.FunctionGraph3D, x, y float64,
		bottom, top []int, bottomColor, topColor color.Color) {
	z := f(x, y)
	xx, yy := fg.GetXX(x, y, z), fg.GetYY(x, y, z)
	if yy > bottom[xx] {
		bottom[xx] = yy
		img.Set(xx, yy, bottomColor)
	}
	if yy < top[xx] {
		top[xx] = yy
		img.Set(xx, yy, topColor)
	}
}
