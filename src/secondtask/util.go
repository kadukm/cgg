package main

import (
	"image"
	"math"

	"github.com/kadukm/cgg/src/utility"
)

func cartesianXToScreen(x float64) int {
	return int((x - xMin) * width / (xMax - xMin))
}

func cartesianYToScreen(y float64) int {
	return int((y - yMax) * height / (yMin - yMax))
}

func screenXXToCartesian(xx int) float64 {
	return float64(xx)*(xMax-xMin)/width + xMin
}

func screenYYToCartesian(yy int) float64 {
	return float64(yy)*(yMin-yMax)/height + yMax
}

func getErrorSizeFor(p point) float64 {
	distanceToDirectrixU := math.Abs(p.u - directrixU)
	distanceToFocus := math.Sqrt(math.Pow(p.u-focus.u, 2) + math.Pow(p.v-focus.v, 2))
	return math.Abs(distanceToDirectrixU - distanceToFocus)
}

func pointInsideImage(img image.Image, p point) bool {
	return (img.Bounds().Min.X <= p.xx && p.xx <= img.Bounds().Max.X &&
		img.Bounds().Min.Y <= p.yy && p.yy <= img.Bounds().Max.Y)
}

//TODO: DRY

func getBestNotUsedPoint(p point, visited map[utility.IntTuple]bool) (res point) {
	minErrorSize := math.MaxFloat64
	for _, currentPoint := range p.ScreenNeighbors() {
		if visited[utility.IntTuple{currentPoint.xx, currentPoint.yy}] {
			continue
		}
		currentErrorSize := getErrorSizeFor(currentPoint)
		if currentErrorSize < minErrorSize {
			minErrorSize = currentErrorSize
			res = currentPoint
		}
	}
	return
}

func getBestPoint(p point) (res point) {
	minErrorSize := math.MaxFloat64
	for _, currentPoint := range p.ScreenNeighbors() {
		currentErrorSize := getErrorSizeFor(currentPoint)
		if currentErrorSize < minErrorSize {
			minErrorSize = currentErrorSize
			res = currentPoint
		}
	}
	return
}

func getBestPointExcluding(p point, excludedPoint point) (res point) {
	minErrorSize := math.MaxFloat64
	for _, currentPoint := range p.ScreenNeighbors() {
		if currentPoint.xx == excludedPoint.xx && currentPoint.yy == excludedPoint.yy {
			continue
		}
		currentErrorSize := getErrorSizeFor(currentPoint)
		if currentErrorSize < minErrorSize {
			minErrorSize = currentErrorSize
			res = currentPoint
		}
	}
	return
}
