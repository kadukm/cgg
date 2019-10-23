package utility

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	// TempFileName is name for file with temp image
	TempFileName = "temp.png"
)

// Point is struct which contains screen coordinates
type Point struct {
	XX, YY int
}

// Rotation returns z-coordinate of vector-product of (p1, p2) * (p2, p3)
func Rotation(p1, p2, p3 Point) int {
	return (p2.XX-p1.XX)*(p3.YY-p2.YY) - (p2.YY-p1.YY)*(p3.XX-p2.XX)
}

// Intersect returns true when segment (p1, p2) intersects with (p3, p4)
func Intersect(p1, p2, p3, p4 Point) bool {
	return Rotation(p1, p2, p3)*Rotation(p1, p2, p4) <= 0 &&
		Rotation(p3, p4, p1)*Rotation(p3, p4, p2) <= 0
}

// EvaluateCos returns cos of angle between vectors (p1, p2) and (p3, p4)
func EvaluateCos(p1, p2, p3, p4 Point) float64 {
	return float64(ScalarProduct(p1, p2, p3, p4)) / (Length(p1, p2) * Length(p3, p4))
}

// ScalarProduct returns scalar product of vectors (p1, p2) and (p3, p4)
func ScalarProduct(p1, p2, p3, p4 Point) int {
	return (p2.XX-p1.XX)*(p4.XX-p3.XX) + (p2.YY-p1.YY)*(p4.YY-p3.YY)
}

// Length returns length of vector (p1, p2)
func Length(p1, p2 Point) float64 {
	dx := p2.XX - p1.XX
	dy := p2.YY - p1.YY
	return math.Sqrt(float64(dx*dx + dy*dy))
}

// PointInsideImage returns true when point (xx, yy) is inside img
func PointInsideImage(img image.Image, xx, yy int) bool {
	return (img.Bounds().Min.X <= xx && xx <= img.Bounds().Max.X &&
		img.Bounds().Min.Y <= yy && yy <= img.Bounds().Max.Y)
}

// SavePNG saves img in PNG format with given filename
func SavePNG(img image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
