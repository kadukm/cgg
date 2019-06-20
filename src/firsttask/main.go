package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/kadukm/cgg/src/utility"
)

const (
	width  = 1000
	height = 600

	a float64 = -5
	b float64 = 13

	minXXAxeStepLength = 32
	minYYAxeStepLength = 32

	notchLength = 5

	fileName = "firsttask.png"
)

var (
	yMin, yMax float64
)

func f(x float64) float64 {
	return math.Cos(x) * x
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	utility.Fill(img, color.White)

	gc := gg.NewContextForImage(img)
	gc.SetLineWidth(1)

	findMinMax()
	drawAxes(gc)
	drawF(gc)

	gc.SavePNG(fileName)
}

func findMinMax() {
	yMin = f(a)
	yMax = yMin

	for xx := 0; xx <= width; xx++ {
		x := screenXXToCartesian(xx)
		y := f(x)
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
	}
}

func drawF(gc *gg.Context) {
	gc.SetColor(color.RGBA{153, 12, 12, 255})

	yy := cartesianYToScreen(f(a))
	gc.MoveTo(0, float64(yy))

	for xx := 1; xx <= width; xx++ {
		x := a + float64(xx)*(b-a)/width
		y := f(x)
		yy = cartesianYToScreen(y)
		gc.LineTo(float64(xx), float64(yy))
	}

	gc.Stroke()
}

func drawAxes(gc *gg.Context) {
	gc.SetColor(color.Black)

	xx0 := cartesianXToScreen(0)
	yy0 := cartesianYToScreen(0)

	gc.DrawLine(float64(xx0), 0, float64(xx0), height)
	gc.DrawLine(0, float64(yy0), width, float64(yy0))

	minXAxeStepLength := screenXXToCartesian(minXXAxeStepLength) - screenXXToCartesian(0)
	xAxeStepLength := math.Max(1, minXAxeStepLength)

	for x := xAxeStepLength; x < math.Max(math.Abs(a), math.Abs(b)); x += xAxeStepLength {
		xRightStr := fmt.Sprintf("%.1f", x)
		xxRight := cartesianXToScreen(x)
		gc.DrawLine(float64(xxRight), float64(yy0-notchLength), float64(xxRight), float64(yy0+notchLength))
		gc.DrawStringAnchored(xRightStr, float64(xxRight), float64(yy0), 0.5, 1)

		xLeftStr := fmt.Sprintf("%.1f", -x)
		xxLeft := cartesianXToScreen(-x)
		gc.DrawLine(float64(xxLeft), float64(yy0-notchLength), float64(xxLeft), float64(yy0+notchLength))
		gc.DrawStringAnchored(xLeftStr, float64(xxLeft), float64(yy0), 0.5, 1)
	}

	minYAxeStepLength := screenYYToCartesian(0) - screenYYToCartesian(minYYAxeStepLength)
	yAxeStepLength := math.Max(1, minYAxeStepLength)

	for y := yAxeStepLength; y < math.Max(math.Abs(yMin), math.Abs(yMax)); y += yAxeStepLength {
		yUpStr := fmt.Sprintf("%.1f", y)
		yyUp := cartesianYToScreen(y)
		gc.DrawLine(float64(xx0-notchLength), float64(yyUp), float64(xx0+notchLength), float64(yyUp))
		gc.DrawStringAnchored(yUpStr, float64(xx0+notchLength), float64(yyUp), 0, 0.5)

		yDownStr := fmt.Sprintf("%.1f", -y)
		yyDown := cartesianYToScreen(-y)
		gc.DrawLine(float64(xx0-notchLength), float64(yyDown), float64(xx0+notchLength), float64(yyDown))
		gc.DrawStringAnchored(yDownStr, float64(xx0), float64(yyDown), 1, 0.5)
	}

	gc.Stroke()
}
