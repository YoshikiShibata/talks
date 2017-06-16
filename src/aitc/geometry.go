package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 } // HL

func Distance(p, q Point) float64 { // 関数 // HL
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 { // メソッド // HL
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 関数呼び出し // HL
	fmt.Println(p.Distance(q))  // メソッド呼び出し // HL
}
