package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/kadukm/cgg/src/utility"
)

const (
	a, b, c float64 = -6, 8, -1

	p = (a + b) / (2 * c)

	xMin, xMax float64 = -10, 10
	yMin, yMax float64 = -6, 6

	width  = 1000
	height = 600

	xAxeStepLength, yAxeStepLength float64 = 1, 1

	notchLength int = 5

	filename = "secondtask.png"
)

var (
	vertex     point
	focus      point
	directrixU float64
)

func init() {
	vertex = pointFromTransformed(0, 0)
	focus = pointFromTransformed(p/2, 0)
	directrixU = -p / 2
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	utility.Fill(img, color.White)

	drawAxes(img)
	drawF(img)

	utility.SavePNG(img, filename)
}

func drawF(img draw.Image) {
	drawingColor := color.RGBA{153, 12, 12, 255}

	img.Set(vertex.xx, vertex.yy, drawingColor)

	visited := make(map[utility.IntTuple]bool)
	visited[utility.IntTuple{vertex.xx, vertex.yy}] = true

	drawParabolaBranch(img, visited, drawingColor)
	drawParabolaBranch(img, visited, drawingColor)
}

func drawParabolaBranch(img draw.Image, visited map[utility.IntTuple]bool, drawingColor color.Color) {
	lastDrawnPoint := vertex
	for pointInsideImage(img, lastDrawnPoint) {
		newPoint := getBestNotUsedPoint(lastDrawnPoint, visited)
		visited[utility.IntTuple{newPoint.xx, newPoint.yy}] = true
		img.Set(newPoint.xx, newPoint.yy, drawingColor)
		lastDrawnPoint = newPoint
	}
}

func drawAxes(img draw.Image) {
	drawingColor := color.Black

	zeroPoint := pointFromCartesian(0, 0)

	utility.DrawVerticalLine(img, zeroPoint.xx, 0, height, drawingColor)
	utility.DrawHorizontalLine(img, zeroPoint.yy, 0, width, drawingColor)

	for x := xAxeStepLength; x < math.Max(math.Abs(xMin), math.Abs(xMax)); x += xAxeStepLength {
		xxRight := cartesianXToScreen(x)
		utility.DrawVerticalLine(img, xxRight, zeroPoint.yy-notchLength, zeroPoint.yy+notchLength, drawingColor)

		xxLeft := cartesianXToScreen(-x)
		utility.DrawVerticalLine(img, xxLeft, zeroPoint.yy-notchLength, zeroPoint.yy+notchLength, drawingColor)
	}

	for y := yAxeStepLength; y < math.Max(math.Abs(yMin), math.Abs(yMax)); y += yAxeStepLength {
		yyUp := cartesianYToScreen(y)
		utility.DrawHorizontalLine(img, yyUp, zeroPoint.xx-notchLength, zeroPoint.xx+notchLength, drawingColor)

		yyDown := cartesianYToScreen(-y)
		utility.DrawHorizontalLine(img, yyDown, zeroPoint.xx-notchLength, zeroPoint.xx+notchLength, drawingColor)
	}
}
