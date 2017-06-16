package main

import "fmt"

func main() {
	var x int = 666

	fmt.Printf("(1)  x = %d\n", x)

	var p *int = &x                 // HL
	fmt.Printf("(2)  p = %p\n", p)  // HL
	fmt.Printf("(3) *p = %d\n", *p) // HL

	*p = 777                       // HL
	fmt.Printf("(4)  x = %d\n", x) // HL
}
