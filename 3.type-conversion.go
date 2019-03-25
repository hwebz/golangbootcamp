package main

import "fmt"

func main() {
	// var i int = 42
	// var f float64 = float64(i)
	// var u uint = uint(f)
	
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Printf("%T(%v) %T(%v) %T(%v)", i, i, f, f, u, u)
}