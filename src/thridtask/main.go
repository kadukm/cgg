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
	triangles := triangulate(p)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	utility.Fill(img, color.White)
	drawTriangulation(img, triangles, color.Black)

	utility.SavePNG(img, filename)
}

func getPolygon() utility.Polygon {
	// Clock-wise direction
	return utility.CreatePolygon([]utility.Point{
		utility.Point{XX: 100, YY: 100},
		utility.Point{XX: 300, YY: 200},
		utility.Point{XX: 400, YY: 100},
		utility.Point{XX: 500, YY: 300},
		utility.Point{XX: 700, YY: 150},
		utility.Point{XX: 800, YY: 300},
		utility.Point{XX: 900, YY: 100},
		utility.Point{XX: 900, YY: 500},
		utility.Point{XX: 100, YY: 500},
		utility.Point{XX: 800, YY: 400},
		utility.Point{XX: 700, YY: 300},
		utility.Point{XX: 400, YY: 400},
		utility.Point{XX: 100, YY: 400},
		utility.Point{XX: 200, YY: 300},
	})

	// Counterlock-wise direction
	// return utility.CreatePolygon([]utility.Point{
	// 	utility.Point{XX: 100, YY: 500},
	// 	utility.Point{XX: 900, YY: 500},
	// 	utility.Point{XX: 500, YY: 400},
	// 	utility.Point{XX: 900, YY: 100},
	// 	utility.Point{XX: 100, YY: 100},
	// })
}

func triangulate(p utility.Polygon) []utility.Triangle {
	nonConvexPointIdx, err := p.TryGetNonConvexPointIdx()
	if err != nil {
		return triangulateConvexPolygon(p)
	}

	goodPointIdx := p.GetGoodPointIdx(nonConvexPointIdx)
	p1, p2 := p.DivideBySegment(nonConvexPointIdx, goodPointIdx)
	triangles1 := triangulate(p1)
	triangles2 := triangulate(p2)

	return append(triangles1, triangles2...)
}

func triangulateConvexPolygon(p utility.Polygon) []utility.Triangle {
	trianglesCount := p.GetPointsCount() - 2
	res := make([]utility.Triangle, 0, trianglesCount)
	rootIdx := 0
	rootPoint := p.GetPointAt(rootIdx)
	curIdx := rootIdx + 1
	for curIdx < p.GetPointsCount()-1 {
		nextIdx := curIdx + 1
		curPoint := p.GetPointAt(curIdx)
		nextPoint := p.GetPointAt(nextIdx)
		newTriangle := utility.Triangle{[3]utility.Point{rootPoint, curPoint, nextPoint}}
		res = append(res, newTriangle)

		curIdx = nextIdx
	}
	return res
}

func drawTriangulation(img draw.Image, ts []utility.Triangle, c color.Color) {
	// TODO: don't draw the same lines of different triangles twice
	for _, triangle := range ts {
		drawTriangle(img, triangle, c)
	}
}

func drawTriangle(img draw.Image, t utility.Triangle, c color.Color) {
	p0 := t.Points[0]
	p1 := t.Points[1]
	p2 := t.Points[2]
	utility.DrawLine(img, p0.XX, p0.YY, p1.XX, p1.YY, c)
	utility.DrawLine(img, p1.XX, p1.YY, p2.XX, p2.YY, c)
	utility.DrawLine(img, p2.XX, p2.YY, p0.XX, p0.YY, c)
}
