package main

import (
	"fmt"
	"unsafe"
)

var global *int

func main() {
	var x int // HL
	xaddr := uintptr(unsafe.Pointer(&x))
	fmt.Printf("x = %#x\n", xaddr)

	y := new(int) // HL
	*y = 1
	yaddr := uintptr(unsafe.Pointer(y))
	fmt.Printf("y = %#x\n", yaddr)

	var z int // HL
	z = 1
	global = &z // HL
	fmt.Printf("z = %p\n", global)
}
