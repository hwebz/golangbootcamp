package main

import (
	"fmt"
	"math"
)

type User struct {
	FirstName, LastName string
}

type Vertex struct {
	X, Y float64
}

// method receicer
func (u *User) Greetings() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// method receiver pass by reference, you can modify the value that its receiver points to
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(u.Greetings())

	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
