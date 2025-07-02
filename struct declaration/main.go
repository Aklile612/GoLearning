package main

import "fmt"

type address struct {
	city     string
	location string
	zipCode  int
}

func main() {
	var a address

	fmt.Println(a)
}
