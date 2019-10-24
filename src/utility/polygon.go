package utility

import (
	"errors"
	"math"
)

func getTheMostLeftPoint(points []Point) (res int) {
	minXX := math.MaxInt64
	for idx, point := range points {
		if point.XX < minXX {
			res = idx
			minXX = point.XX
		}
	}
	return res
}

func getClockWiseDelta(curPoint, prevNeighbor, nextNeighbor Point) int {
	if (EvaluateCos(Point{0, 1}, Point{0, 0}, curPoint, prevNeighbor) >
		EvaluateCos(Point{0, 1}, Point{0, 0}, curPoint, nextNeighbor)) {
		return -1
	}
	return 1
}

func normalizeIdx(rawIdx, count int) int {
	res := rawIdx % count
	if res < 0 {
		res = res + count
	}
	return res
}

func CreatePolygon(points []Point) Polygon {
	leftPointIdx := getTheMostLeftPoint(points)
	prevNeighborIdx := normalizeIdx(leftPointIdx-1, len(points))
	nextNeighborIdx := normalizeIdx(leftPointIdx+1, len(points))
	clockWiseDelta := getClockWiseDelta(
		points[leftPointIdx], points[prevNeighborIdx], points[nextNeighborIdx])

	return Polygon{clockWiseDelta, points}
}

type Polygon struct {
	ClockWiseDelta int
	points         []Point
}

func (p Polygon) GetPointsCount() int {
	return len(p.points)
}

func (p Polygon) GetPointAt(idx int) Point {
	normalizedIdx := p.normalizeIdx(idx)
	return p.points[normalizedIdx]
}

func (p Polygon) normalizeIdx(rawIdx int) int {
	return normalizeIdx(rawIdx, len(p.points))
}

func (p Polygon) TryGetNonConvexPointIdx() (int, error) {
	const startIdx = 0
	idx1 := startIdx
	for i := 0; i < p.GetPointsCount(); i++ {
		idx2 := idx1 + p.ClockWiseDelta
		newPointIdx := idx2 + p.ClockWiseDelta
		p1 := p.GetPointAt(idx1)
		p2 := p.GetPointAt(idx2)
		newPoint := p.GetPointAt(newPointIdx)
		if Rotation(p1, p2, newPoint) < 0 {
			return p.normalizeIdx(idx2), nil
		}
		idx1 = idx2
	}
	return 0, errors.New("can't find non-convex point")

}

func (p Polygon) GetDividingPointIdx(nonConvexPointIdx int) int {
	for pointIdx := 0; pointIdx < len(p.points); pointIdx++ {
		if p.arePointsNeighbors(nonConvexPointIdx, pointIdx) {
			continue
		}
		//TODO: I can handle ends of intersected segments, but now I won't do it
		if p.isSegmentStartsInside(nonConvexPointIdx, pointIdx) &&
			p.isSegmentInside(nonConvexPointIdx, pointIdx) {
			return pointIdx
		}
	}
	panic("impossible situation")
}

func (p Polygon) arePointsNeighbors(idx1, idx2 int) bool {
	if idx1 == idx2 {
		return true
	}

	idx2Neghbor1 := p.normalizeIdx(idx2 - 1)
	idx2Neghbor2 := p.normalizeIdx(idx2 + 1)
	return idx1 == idx2Neghbor1 || idx1 == idx2Neghbor2
}

func (p Polygon) DivideBySegment(idx1, idx2 int) (Polygon, Polygon) {
	p1 := p.GetSubPolygon(idx1, idx2)
	p2 := p.GetSubPolygon(idx2, idx1)
	return p1, p2
}

func (p Polygon) GetSubPolygon(idxFrom, idxTo int) Polygon {
	var neededCapacity int
	if idxFrom <= idxTo {
		neededCapacity = idxTo - idxFrom + 1
	} else {
		neededCapacity = p.GetPointsCount() - idxFrom + idxTo + 1
	}
	points := make([]Point, 0, neededCapacity)
	curIdx := idxFrom
	for curIdx != idxTo {
		points = append(points, p.GetPointAt(curIdx))
		curIdx = p.normalizeIdx(curIdx + 1)
	}
	points = append(points, p.GetPointAt(curIdx))
	return Polygon{p.ClockWiseDelta, points}
}

func (p Polygon) isSegmentStartsInside(pointStartIdx, pointEndIdx int) bool {
	pointStart := p.GetPointAt(pointStartIdx)
	pointEnd := p.GetPointAt(pointEndIdx)

	prevNeighbor := p.GetPointAt(pointStartIdx - p.ClockWiseDelta)
	if Rotation(prevNeighbor, pointStart, pointEnd) >= 0 {
		return true
	}

	nextNeighbor := p.GetPointAt(pointStartIdx + p.ClockWiseDelta)
	if EvaluateCos(prevNeighbor, pointStart, pointStart, nextNeighbor) >
		EvaluateCos(prevNeighbor, pointStart, pointStart, pointEnd) {
		return true
	}

	return false
}

func (p Polygon) isSegmentInside(pointStartIdx, pointEndIdx int) bool {
	pointStart := p.GetPointAt(pointStartIdx)
	pointEnd := p.GetPointAt(pointEndIdx)

	for pointIdx := 0; pointIdx < len(p.points); pointIdx++ {
		if pointIdx == pointStartIdx || pointIdx == pointEndIdx {
			continue
		}
		nextPointIdx := p.normalizeIdx(pointIdx + 1)
		if nextPointIdx == pointStartIdx || nextPointIdx == pointEndIdx {
			continue
		}

		point := p.GetPointAt(pointIdx)
		nextPoint := p.GetPointAt(nextPointIdx)
		if Intersect(pointStart, pointEnd, point, nextPoint) {
			return false
		}
	}

	return true
}
