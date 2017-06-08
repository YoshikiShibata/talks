package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	stackSize := debug.SetMaxStack(64 * 1024)
	fmt.Printf("Previous Stack Size : %d\n", stackSize)
}
