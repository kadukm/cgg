package main

import (
	"github.com/kadukm/cgg/src/utility"
)

// These values allow transfrom points without redundant recount
const (
	t1 float64 = b * b / (4 * c * (a + b))
	t2 float64 = b / (2 * c)
)

type point struct {
	x, y   float64
	u, v   float64
	xx, yy int
}

func pointFromTransformed(fg utility.FunctionGraph2d, u, v float64) point {
	x, y := -u+t1, v-u+t1-t2
	xx, yy := fg.CartesianXToScreen(x), fg.CartesianYToScreen(y)
	return point{x, y, u, v, xx, yy}
}

func pointFromCartesian(fg utility.FunctionGraph2d, x, y float64) point {
	u, v := -x+t1, y-x+t2
	xx, yy := fg.CartesianXToScreen(x), fg.CartesianYToScreen(y)
	return point{x, y, u, v, xx, yy}
}

func pointFromScreen(fg utility.FunctionGraph2d, xx, yy int) point {
	x, y := fg.ScreenXXToCartesian(xx), fg.ScreenYYToCartesian(yy)
	u, v := -x+t1, y-x+t2
	return point{x, y, u, v, xx, yy}
}

func (p *point) ScreenNeighbors(fg utility.FunctionGraph2d) (neighbors [8]point) {
	deltas := []struct{ dxx, dyy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for i, delta := range deltas {
		newXX := p.xx + delta.dxx
		newYY := p.yy + delta.dyy
		neighbor := pointFromScreen(fg, newXX, newYY)
		neighbors[i] = neighbor
	}
	return neighbors
}
