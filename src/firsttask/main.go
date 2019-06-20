package main

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/kadukm/cgg/src/utility"
)

const (
	width  = 1000
	height = 600

	a float64 = -10
	b float64 = 15

	minXXAxeStepLength = 32
	minYYAxeStepLength = 32

	fileName = "firsttask.png"
)

var (
	yMin, yMax float64
)

func f(x float64) float64 {
	return x * x * x
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

	gc.MoveTo(float64(xx0), 0)
	gc.LineTo(float64(xx0), height)

	gc.MoveTo(0, float64(yy0))
	gc.LineTo(width, float64(yy0))

	// minXAxeStepLength := screenXXToCartesian(minXXAxeStepLength) - screenXXToCartesian(0)
	// xAxeStepLength := math.Max(1, minXAxeStepLength)

	// for x := float64(0); x < math.Max(math.Abs(a), math.Abs(b)); x += xAxeStepLength {
	// 	xx_right := cartesianXToScreen(x)
	// 	xx_left := cartesianXToScreen(-x)
	// 	qp.drawLine(xx_right, yy0 - 2, xx_right, yy0 + 2)
	// 	qp.drawLine(xx_left, yy0 - 2, xx_left, yy0 + 2)
	// 	qp.drawText(QRect(xx_right - 10, yy0 + 10, 20, 20), Qt.AlignCenter, str(x))
	// 	qp.drawText(QRect(xx_left - 10, yy0 + 10, 20, 20), Qt.AlignCenter, str(-x))
	// }

	gc.Stroke()

}
