package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/kadukm/cgg/src/utility"
)

const (
	a, b, c float64 = -6, 8, -1

	xMin, xMax float64 = -10, 10
	yMin, yMax float64 = -6, 6

	width  = 1000
	height = 600

	xAxeStepLength float64 = 1
	yAxeStepLength float64 = 1

	notchLength int = 5

	filename = "secondtask.png"
)

func main() {
	fg := initFunctionGraph()
	po := initParabolaOptions(fg)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	utility.DrawAxes(img, fg, color.Black)
	drawFunction(img, fg, po)

	utility.SavePNG(img, filename)
}

func initFunctionGraph() (fg utility.FunctionGraph2d) {
	fg.XMin, fg.XMax = xMin, xMax
	fg.YMin, fg.YMax = yMin, yMax
	fg.Width = width
	fg.Height = height

	fg.XAxeStepLength = xAxeStepLength
	fg.YAxeStepLength = yAxeStepLength
	fg.NotchLength = notchLength

	return
}

func initParabolaOptions(fg utility.FunctionGraph2d) (po parabolaOptions) {
	po.p = (a + b) / (2 * c)
	po.focus = pointFromTransformed(fg, po.p/2, 0)
	po.directrixU = -po.p / 2
	po.vertex = pointFromTransformed(fg, 0, 0)

	return
}

func drawFunction(img draw.Image, fg utility.FunctionGraph2d, po parabolaOptions) {
	drawingColor := color.RGBA{153, 12, 12, 255}

	img.Set(po.vertex.xx, po.vertex.yy, drawingColor)

	visited := make(map[utility.Point]bool)
	visited[utility.Point{po.vertex.xx, po.vertex.yy}] = true

	drawParabolaBranch(img, fg, po, visited, drawingColor)
	drawParabolaBranch(img, fg, po, visited, drawingColor)
}

func drawParabolaBranch(
	img draw.Image,
	fg utility.FunctionGraph2d,
	po parabolaOptions,
	visited map[utility.Point]bool,
	drawingColor color.Color,
) {
	lastDrawnPoint := po.vertex
	for utility.PointInsideImage(img, lastDrawnPoint.xx, lastDrawnPoint.yy) {
		newPoint := getNearestNotUsedPoint(lastDrawnPoint, visited, fg, po)
		visited[utility.Point{newPoint.xx, newPoint.yy}] = true
		img.Set(newPoint.xx, newPoint.yy, drawingColor)
		lastDrawnPoint = newPoint
	}
}
