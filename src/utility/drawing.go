package utility

import (
	"image/color"
	"image/draw"
	"math"
)

//Fill fills img with color c
func Fill(img draw.Image, c color.Color) {
	for xx := img.Bounds().Min.X; xx <= img.Bounds().Max.X; xx++ {
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

//TODO: refactor it
func DrawLine(img draw.Image, xx1, yy1, xx2, yy2 int, color color.Color) {
	var dx, dy int
	var sx, sy int

	if xx1 < xx2 {
		dx = xx2 - xx1
		sx = 1
	} else {
		dx = xx1 - xx2
		sx = -1
	}

	if yy1 < yy2 {
		dy = yy2 - yy1
		sy = 1
	} else {
		dy = yy1 - yy2
		sy = -1
	}

	err := dx - dy

	for {
		img.Set(xx1, yy1, color)
		if xx1 == xx2 && yy1 == yy2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			xx1 += sx
		}
		if e2 < dx {
			err += dx
			yy1 += sy
		}
	}
}

//DrawAxes draws axes with branches for fg
func DrawAxes(img draw.Image, fg FunctionGraph, color color.Color) {
	xx0, yy0 := fg.CartesianXToScreen(0), fg.CartesianYToScreen(0)

	DrawVerticalLine(img, xx0, 0, fg.Height, color)
	DrawHorizontalLine(img, yy0, 0, fg.Width, color)

	for x := fg.XAxeStepLength; x < math.Max(math.Abs(fg.XMin), math.Abs(fg.XMax)); x += fg.XAxeStepLength {
		xxRight := fg.CartesianXToScreen(x)
		DrawVerticalLine(img, xxRight, yy0-fg.NotchLength, yy0+fg.NotchLength, color)

		xxLeft := fg.CartesianXToScreen(-x)
		DrawVerticalLine(img, xxLeft, yy0-fg.NotchLength, yy0+fg.NotchLength, color)
	}

	for y := fg.YAxeStepLength; y < math.Max(math.Abs(fg.YMin), math.Abs(fg.YMax)); y += fg.YAxeStepLength {
		yyUp := fg.CartesianYToScreen(y)
		DrawHorizontalLine(img, yyUp, xx0-fg.NotchLength, xx0+fg.NotchLength, color)

		yyDown := fg.CartesianYToScreen(-y)
		DrawHorizontalLine(img, yyDown, xx0-fg.NotchLength, xx0+fg.NotchLength, color)
	}
}
