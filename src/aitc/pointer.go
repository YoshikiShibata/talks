package main

import "fmt"

func main() {
	var x int = 666

	fmt.Printf("(1)  x = %d\n", x)

	var p *int = &x
	fmt.Printf("(2)  p = %p\n", p)
	fmt.Printf("(3) *p = %d\n", *p)

	*p = 777
	fmt.Printf("(4)  x = %d\n", x)
}
