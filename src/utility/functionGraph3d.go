package utility

import "math"

var (
	dimetricValue  = 2 * math.Sqrt2
	isometricValue = math.Sqrt(3) / 2
)

type FunctionGraph3d struct {
	XMin, XMax float64
	YMin, YMax float64

	Width  int
	Height int

	LinesCount       int
	StepsByLineCount int

	XXProjectionMin, XXProjectionMax float64
	YYProjectionMin, YYProjectionMax float64
}

func (fg FunctionGraph3d) GetXX(x, y, z float64) int {
	xxProjection := fg.GetXXProjection(x, y, z)
	return fg.XXProjectionToScreen(xxProjection)
}

func (fg FunctionGraph3d) GetYY(x, y, z float64) int {
	yyProjection := fg.GetYYProjection(x, y, z)
	return fg.YYProjectionToScreen(yyProjection)
}

func (fg FunctionGraph3d) XXProjectionToScreen(xxProjection float64) int {
	return int((xxProjection - fg.XXProjectionMin) * float64(fg.Width) / (fg.XXProjectionMax - fg.XXProjectionMin))
}

func (fg FunctionGraph3d) YYProjectionToScreen(yyProjection float64) int {
	return int((yyProjection - fg.YYProjectionMin) * float64(fg.Height) / (fg.YYProjectionMax - fg.YYProjectionMin))
}

func (fg FunctionGraph3d) GetXXProjection(x, y, z float64) float64 {
	//return fg.getXXDimetricProjection(x, y)
	return fg.getXXIsometricProjection(x, y)
}

func (fg FunctionGraph3d) GetYYProjection(x, y, z float64) float64 {
	//return fg.getYYDimetricProjection(x, z)
	return fg.getYYIsometricProjection(x, y, z)
}

func (fg FunctionGraph3d) getXXDimetricProjection(x, y float64) float64 {
	return y - x/dimetricValue
}

func (fg FunctionGraph3d) getYYDimetricProjection(x, z float64) float64 {
	return x/dimetricValue - z
}

func (fg FunctionGraph3d) getXXIsometricProjection(x, y float64) float64 {
	return (y - x) * isometricValue
}

func (fg FunctionGraph3d) getYYIsometricProjection(x, y, z float64) float64 {
	return (x+y)/2 - z
}
