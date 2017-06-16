// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 161.

// Coloredpoint demonstrates struct embedding.
package main

import (
	"fmt"
	"math"
)

// START OMIT
import "image/color"

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point // 埋め込み // HL
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// END OMIT

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point)) // "5" // HL
	p.ScaleBy(2)                     // HL
	q.ScaleBy(2)                     // HL
	fmt.Println(p.Distance(q.Point)) // "10" // HL
}
