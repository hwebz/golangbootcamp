package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://hawebz.com")
	fmt.Println(resp)
	fmt.Println(err)
}