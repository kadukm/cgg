package main

import (
	"math"

	"github.com/kadukm/cgg/src/utility"
)

type parabolaOptions struct {
	p          float64
	directrixU float64
	vertex     point
	focus      point
}

func (po parabolaOptions) getErrorSizeFor(p point) float64 {
	distanceToDirectrixU := math.Abs(p.u - po.directrixU)
	distanceToFocus := math.Sqrt(math.Pow(p.u-po.focus.u, 2) + math.Pow(p.v-po.focus.v, 2))
	return math.Abs(distanceToDirectrixU - distanceToFocus)
}

func getNearestNotUsedPoint(
	p point,
	visited map[utility.Point]bool,
	fg utility.FunctionGraph2D,
	po parabolaOptions,
) (res point) {
	minErrorSize := math.MaxFloat64
	for _, currentPoint := range p.ScreenNeighbors(fg) {
		if visited[utility.Point{currentPoint.xx, currentPoint.yy}] {
			continue
		}
		currentErrorSize := po.getErrorSizeFor(currentPoint)
		if currentErrorSize < minErrorSize {
			minErrorSize = currentErrorSize
			res = currentPoint
		}
	}
	return
}
