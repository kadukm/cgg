package main

func cartesianXToScreen(x float64) int {
	return int((x - a) * width / (b - a))
}

func cartesianYToScreen(y float64) int {
	return int((y - yMax) * height / (yMin - yMax))
}

func screenXXToCartesian(xx int) float64 {
	return float64(xx)*(b-a)/width + a
}

func screenYYToCartesian(yy int) float64 {
	return float64(yy)*(yMin-yMax)/height + yMax
}
