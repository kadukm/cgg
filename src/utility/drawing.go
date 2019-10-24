package utility

import (
	"image/color"
	"image/draw"
	"math"
)

//Fill fills img with color c
func Fill(img draw.Image, c color.Color) {
	for xx := img.Bounds().Min.X; xx < img.Bounds().Max.X; xx++ {
		for yy := img.Bounds().Min.Y; yy < img.Bounds().Max.Y; yy++ {
			img.Set(xx, yy, c)
		}
	}
}

//DrawHorizontalLine draws line from (xxStart, yy0) to (xxEnd, yy0)
func DrawHorizontalLine(img draw.Image, yy0, xxStart, xxEnd int, c color.Color) {
	for xx := xxStart; xx <= xxEnd; xx++ {
		img.Set(xx, yy0, c)
	}
}

//DrawVerticalLine draws line from (xx0, yyStart) to (xx0, yyEnd)
func DrawVerticalLine(img draw.Image, xx0, yyStart, yyEnd int, c color.Color) {
	for yy := yyStart; yy < yyEnd; yy++ {
		img.Set(xx0, yy, c)
	}
}

//DrawLine draws line from (xx1, yy1) to (xx2, yy2). Coordinates can be any
func DrawLine(img draw.Image, xx1, yy1, xx2, yy2 int, color color.Color) {
	var dx, dy int
	var xxStep, yyStep int

	if xx1 < xx2 {
		dx = xx2 - xx1
		xxStep = 1
	} else {
		dx = xx1 - xx2
		xxStep = -1
	}

	if yy1 < yy2 {
		dy = yy2 - yy1
		yyStep = 1
	} else {
		dy = yy1 - yy2
		yyStep = -1
	}

	err := dx - dy

	img.Set(xx1, yy1, color)

	xx, yy := xx1, yy1
	for xx != xx2 || yy != yy2 {
		currentErr := 2 * err
		if currentErr > -dy {
			err -= dy
			xx += xxStep
		}
		if currentErr < dx {
			err += dx
			yy += yyStep
		}

		img.Set(xx, yy, color)
	}
}

//DrawAxes draws axes with branches for fg
func DrawAxes(img draw.Image, fg FunctionGraph2d, c color.Color) {
	xx0, yy0 := fg.CartesianXToScreen(0), fg.CartesianYToScreen(0)

	DrawVerticalLine(img, xx0, 0, fg.Height, c)
	DrawHorizontalLine(img, yy0, 0, fg.Width, c)

	for x := fg.XAxeStepLength; x < math.Max(math.Abs(fg.XMin), math.Abs(fg.XMax)); x += fg.XAxeStepLength {
		xxRight := fg.CartesianXToScreen(x)
		DrawVerticalLine(img, xxRight, yy0-fg.NotchLength, yy0+fg.NotchLength, c)

		xxLeft := fg.CartesianXToScreen(-x)
		DrawVerticalLine(img, xxLeft, yy0-fg.NotchLength, yy0+fg.NotchLength, c)
	}

	for y := fg.YAxeStepLength; y < math.Max(math.Abs(fg.YMin), math.Abs(fg.YMax)); y += fg.YAxeStepLength {
		yyUp := fg.CartesianYToScreen(y)
		DrawHorizontalLine(img, yyUp, xx0-fg.NotchLength, xx0+fg.NotchLength, c)

		yyDown := fg.CartesianYToScreen(-y)
		DrawHorizontalLine(img, yyDown, xx0-fg.NotchLength, xx0+fg.NotchLength, c)
	}
}
