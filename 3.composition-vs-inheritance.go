package main

import "fmt"

type User struct {
	Id int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	// Id int
	// Name, Location string
	//User // pass by value
	*User // pass by reference
	GameId int
}

func main() {
	// p := Player{}
	// p.Id = 42
	// p.Name = "Matt"
	// p.Location = "LA"
	// p.GameId = 90404
	// fmt.Printf("%+v", p)

	p := Player {
		//User{Id: 42, Name: "Matt", Location: "LA"},
		&User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}
	fmt.Printf("Id: %d, Name: %s, Location: %s, Game id: %d\n", p.Id, p.Name, p.Location, p.GameId)
	fmt.Println(p.Greetings())

	// Directly set a field defined on the Player struct
	p.Id = 11
	fmt.Printf("%+v", p)
}