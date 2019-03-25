package main

import (
	"fmt"
	"math/cmplx"
)

var (
	GoIsFun bool = true
	maxInt uint64 = 1 << 64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, GoIsFun, GoIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}