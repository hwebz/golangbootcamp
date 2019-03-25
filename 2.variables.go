package main

import "fmt"

func main() {
	// var (
	// 	name		string
	// 	age			int
	// 	location	string
	// )

	// var (
	// 	name, location string
	// 	age int
	// )

	// var name string
	// var age int
	// var location string

	// with initialization value
	// var (
	// 	name		string = "Prince Oberyn"
	// 	age			int = 32
	// 	location	string = "Dorne"
	// )

	var (
		name, location, age = "Prince Oberyn", "Dorne", 32
	)

	// var (
	// 	name, location string = "Prince Oberyn, Dorne"
	// 	age int = 32
	// )

	// var name string = "Prince Oberyn"
	// var age int = 32
	// var location string = "Dorne"

	// name, location := "Prince Oberyn", "Dorne"
	// age := 32

	fmt.Printf("%s (%d) %s", name, age, location)

	action := func() {
		// do something
	}
	action()
}