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

	linesCount = 50
	stepsCount = width * 2

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

func initFunctionGraph() (fg utility.FunctionGraph3d) {
	fg.XMin, fg.XMax = x1, x2
	fg.YMin, fg.YMax = y1, y2
	fg.Width = width
	fg.Height = height
	fg.LinesCount = linesCount
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

	prevBottomMax, prevTopMin := drawFunctionMovingByX(img, fg, bottomColor, topColor)
	drawFunctionMovingByY(img, fg, prevBottomMax, prevTopMin, bottomColor, topColor)
}

func drawFunctionMovingByX(img draw.Image, fg utility.FunctionGraph3d,
		bottomColor color.Color, topColor color.Color) ([]int, []int) {
	bottom := utility.InitIntSlice(fg.Width+1, -1)
	bottomMax := utility.InitIntSlice(fg.Width+1, fg.Height+1)
	top := utility.InitIntSlice(fg.Width+1, fg.Height+1)
	topMin := utility.InitIntSlice(fg.Width+1, -1)
	for i := 0; i <= fg.LinesCount; i++ {
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

			if yy < bottomMax[xx] && img.At(xx, yy) == bottomColor {
				bottomMax[xx] = yy
			} else if yy > topMin[xx] && img.At(xx, yy) == topColor {
				topMin[xx] = yy
			}
		}
	}

	return bottomMax, topMin
}

func drawFunctionMovingByY(img draw.Image, fg utility.FunctionGraph3d,
		prevBottomMax []int, prevTopMin []int,
		bottomColor color.Color, topColor color.Color) {
	bottom := utility.InitIntSlice(fg.Width+1, -1)
	top := utility.InitIntSlice(fg.Width+1, fg.Height+1)
	for j := 0; j <= fg.LinesCount; j++ {
		y := fg.YMax - (fg.YMax-fg.YMin)*float64(j)/float64(fg.LinesCount)
		for i := 0; i <= fg.StepsCount; i++ {
			x := fg.XMax - (fg.XMax-fg.XMin)*float64(i)/float64(fg.StepsCount)
			z := f(x, y)

			xxProjection := fg.GetXXProjection(x, y, z)
			yyProjection := fg.GetYYProjection(x, y, z)

			xx := fg.XXProjectionToScreen(xxProjection)
			yy := fg.YYProjectionToScreen(yyProjection)
			if yy > bottom[xx] && yy > prevTopMin[xx] {
				bottom[xx] = yy
				img.Set(xx, yy, bottomColor)
			}
			if yy < top[xx] && yy < prevBottomMax[xx] {
				top[xx] = yy
				img.Set(xx, yy, topColor)
			}
		}
	}
}
