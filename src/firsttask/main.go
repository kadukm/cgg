package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/kadukm/cgg/src/utility"
)

const (
	a, b float64 = -5, 13

	width  = 1000
	height = 600

	xAxeStepLength float64 = 1
	yAxeStepLength float64 = 1

	notchLength = 5

	fileName = "firsttask.png"
)

func f(x float64) float64 {
	return math.Cos(x) * x
}

func main() {
	fg := initFunctionGraph()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	utility.DrawAxes(img, fg, color.Black)
	drawFunction(img, fg)

	utility.SavePNG(img, fileName)
}

func initFunctionGraph() (fg utility.FunctionGraph2d) {
	fg.XMin, fg.XMax = a, b
	fg.Width = width - 1
	fg.Height = height - 1

	findMinMaxY(&fg)

	fg.XAxeStepLength = xAxeStepLength
	fg.YAxeStepLength = yAxeStepLength
	fg.NotchLength = notchLength

	return
}

func findMinMaxY(fg *utility.FunctionGraph2d) {
	tempValue := f(a)
	fg.YMin = tempValue
	fg.YMax = tempValue

	for xx := 0; xx <= width; xx++ {
		x := fg.ScreenXXToCartesian(xx)
		y := f(x)
		if y < fg.YMin {
			fg.YMin = y
		}
		if y > fg.YMax {
			fg.YMax = y
		}
	}
}

func drawFunction(img draw.Image, fg utility.FunctionGraph2d) {
	drawingColor := color.RGBA{R: 153, G: 12, B: 12, A: 255}

	prevYY := fg.CartesianYToScreen(f(a))
	prevXX := 0
	for newXX := 1; newXX <= width; newXX++ {
		x := a + float64(newXX)*(b-a)/width
		y := f(x)
		newYY := fg.CartesianYToScreen(y)
		utility.DrawLine(img, prevXX, prevYY, newXX, newYY, drawingColor)
		prevXX, prevYY = newXX, newYY
	}
}
