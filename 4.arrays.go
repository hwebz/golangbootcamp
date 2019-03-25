package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "Error"
	// fmt.Println(a[0], a[1]) // out of bounds error
	fmt.Println(a)

	b := [2]string{"hello", "world!"}
	fmt.Printf("%q", b)

	c := [...]string{"hello", "world"}
	fmt.Printf("%q", c)

	fmt.Printf("%s\n", a)

	var d [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = fmt.Sprintf("row %d - column %d", i + 1, j + 1)
		}
	}
	fmt.Printf("%q", d)
}