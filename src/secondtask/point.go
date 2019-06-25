package main

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

func pointFromTransformed(u, v float64) point {
	x, y := -u+t1, v-u+t1-t2
	xx, yy := cartesianXToScreen(x), cartesianYToScreen(y)
	return point{x, y, u, v, xx, yy}
}

func pointFromCartesian(x, y float64) point {
	u, v := -x+t1, y-x+t2
	xx, yy := cartesianXToScreen(x), cartesianYToScreen(y)
	return point{x, y, u, v, xx, yy}
}

func pointFromScreen(xx, yy int) point {
	x, y := screenXXToCartesian(xx), screenYYToCartesian(yy)
	u, v := -x+t1, y-x+t2
	return point{x, y, u, v, xx, yy}
}

func (p *point) ScreenNeighbors() (neighbors [8]point) {
	deltas := []struct{ dxx, dyy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for i, delta := range deltas {
		newXX := p.xx + delta.dxx
		newYY := p.yy + delta.dyy
		neighbor := pointFromScreen(newXX, newYY)
		neighbors[i] = neighbor
	}
	return neighbors
}
