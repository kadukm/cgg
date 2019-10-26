package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/kadukm/cgg/src/utility"
)

const (
	width  = 1000
	height = 600

	filename = "thirdtask.png"
)

func main() {
	p := getPolygon()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	drawPolygon(img, p, color.Black)
	drawTriangulation(img, p, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	utility.SavePNG(img, filename)
}

func getPolygon() utility.Polygon {
	// Clock-wise direction
	return utility.CreatePolygon([]utility.Point{
		{XX: 100, YY: 100},
		{XX: 300, YY: 200},
		{XX: 400, YY: 100},
		{XX: 500, YY: 300},
		{XX: 700, YY: 150},
		{XX: 800, YY: 300},
		{XX: 900, YY: 100},
		{XX: 900, YY: 500},
		{XX: 100, YY: 500},
		{XX: 800, YY: 400},
		{XX: 700, YY: 300},
		{XX: 400, YY: 400},
		{XX: 100, YY: 400},
		{XX: 200, YY: 300},
	})
	//Counterlock-wise direction
	//return utility.CreatePolygon([]utility.Point{
	//	{XX: 100, YY: 500},
	//	{XX: 900, YY: 500},
	//	{XX: 500, YY: 400},
	//	{XX: 900, YY: 100},
	//	{XX: 100, YY: 100},
	//})
}

func drawPolygon(img draw.Image, p utility.Polygon, c color.Color) {
	for curIdx := 0; curIdx < p.GetPointsCount(); curIdx++ {
		curPoint := p.GetPointAt(curIdx)
		nextPoint := p.GetPointAt(curIdx + 1)
		utility.DrawLine(img, curPoint.XX, curPoint.YY, nextPoint.XX, nextPoint.YY, c)
	}
}

func drawTriangulation(img draw.Image, p utility.Polygon, c color.Color) {
	nonConvexPointIdx, err := p.TryGetNonConvexPointIdx()
	if err != nil {
		drawTriangulationOfConvexPolygon(img, p, c)
		return
	}

	dividingPointIdx := p.GetDividingPointIdx(nonConvexPointIdx)

	nonConvexPoint := p.GetPointAt(nonConvexPointIdx)
	dividingPoint := p.GetPointAt(dividingPointIdx)
	utility.DrawLine(img, nonConvexPoint.XX, nonConvexPoint.YY, dividingPoint.XX, dividingPoint.YY, c)

	p1, p2 := p.DivideBySegment(nonConvexPointIdx, dividingPointIdx)
	drawTriangulation(img, p1, c)
	drawTriangulation(img, p2, c)
}

func drawTriangulationOfConvexPolygon(img draw.Image, p utility.Polygon, c color.Color) {
	rootIdx := 0
	rootPoint := p.GetPointAt(rootIdx)
	for curIdx := rootIdx + 2; curIdx < p.GetPointsCount()-1; curIdx++ {
		curPoint := p.GetPointAt(curIdx)
		utility.DrawLine(img, rootPoint.XX, rootPoint.YY, curPoint.XX, curPoint.YY, c)
	}
}
