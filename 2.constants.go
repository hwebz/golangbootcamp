package main

// import "fmt"
// import "math/rand"
import (
	"fmt"
)

const Pi = 3.14
const (
	StatusOK					= 200
	StatusCreated				= 201
	StatusAccepted				= 202
	StatusNonAuthoritativeInfo 	= 203
	StatusNoContent				= 204
	StatusResetContent			= 205
	StatusPartialContent		= 206
)
const (
	Truth = false
	Big = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)

	cylonModel := 6
	fmt.Println(cylonModel)

	name := "Caprica-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s", name, aka)
}