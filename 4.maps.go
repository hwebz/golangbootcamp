package main

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

var m = map[string]Vertex{
	"Google": {37.435454, -122.32434},
}

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt", "Emma", "Isabella", "Emily", "Madison", "Ava", "Olivia", "Sophia", "Abigail", "Elizabeth", "Chloe", "Smantha", "Addison", "Natalie", "Mia", "Alexis"}

func main() {
	cities := map[string]int{
		"New York": 8336697,
		"Los Angeles": 3857799,
		"Chicago": 2714856,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}

	var maxLen int
	for _, name := range names {
		if l := len(name); l > maxLen {
			maxLen = l
		}
	}
	output := make([][]string, maxLen)
	for _, name := range names {
		output[len(name)-1] = append(output[len(name)-1], name)
	}
	fmt.Printf("%v", output)
	fmt.Printf("%#v", cities)

	// m[key] = elem, elem = m[key], delete(m, key), elem, ok = m[key]
	// m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68423, -75.43432}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
}