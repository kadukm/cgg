package utility

type FunctionGraph struct {
	XMin, XMax float64
	YMin, YMax float64

	Width  int
	Height int

	XAxeStepLength float64
	YAxeStepLength float64

	NotchLength int
}

func (fg FunctionGraph) CartesianXToScreen(x float64) int {
	return int((x - fg.XMin) * float64(fg.Width) / (fg.XMax - fg.XMin))
}

func (fg FunctionGraph) CartesianYToScreen(y float64) int {
	return int((y - fg.YMax) * float64(fg.Height) / (fg.YMin - fg.YMax))
}

func (fg FunctionGraph) ScreenXXToCartesian(xx int) float64 {
	return float64(xx)*(fg.XMax-fg.XMin)/float64(fg.Width) + fg.XMin
}

func (fg FunctionGraph) ScreenYYToCartesian(yy int) float64 {
	return float64(yy)*(fg.YMin-fg.YMax)/float64(fg.Height) + fg.YMax
}
