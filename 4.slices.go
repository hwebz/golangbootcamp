package main

import "fmt"

func main() {
	p := []int{2,3,5,7,11,13}
	fmt.Println(p)

	fmt.Println(p[1:4])

	fmt.Println(p[:3])

	fmt.Println(p[4:])

	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Vehicle"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)

	countries := []string{}
	countries = append(countries, "Vietnam")
	countries = append(countries, "San Diego", "Mountain View")
	otherCountries := []string{"Cuba", "Russia"}
	countries = append(countries, otherCountries...)
	fmt.Println(countries)

	fmt.Println(len(countries))

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}