package main

import "fmt"

func main() {

	var myMap = make(map[string]float64)

	myMap["a"] = 98
	myMap["N"] = 86.75
	myMap["Na"] = 80.5
	myMap["ya"] = 79.8

	fmt.Println("The total thing is: ", myMap)
}
