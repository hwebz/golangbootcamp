package main

import (
	"errors"
	"fmt"
)

func foo() error {
	if true {
		return errors.New("Something went wrong")
	}
	return nil
}

func panic(err error) {
	fmt.Println("Error: ", err)
}

func main() {
	answer := 43

	if answer != 42 {
		fmt.Println("Wrong answer")
	}

	if err := foo(); err != nil {
		panic(err)
	}
}
