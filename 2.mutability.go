package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs int
}

func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func newReleaseRef(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

	me2 := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me2.Name, newReleaseRef(me2))
	fmt.Printf("%s has a total of %d songs", me2.Name, me2.Songs)
}