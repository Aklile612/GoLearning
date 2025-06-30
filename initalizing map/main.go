package main

import "fmt"

func main() {

	var emptyMap = map[int]string{}

	fmt.Println("Empty map: ", emptyMap)

	var myMap = map[int]string{
		1:"A",
		2:"B",
		3:"C",
		4:"D",
	}

	fmt.Println("My map: ",myMap)
}
