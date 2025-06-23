package main

import "fmt"

func main() {

	var mainslice = []int{1, 23, 4, 5, 67, 8, 5}

	newslice := mainslice[3:6]

	fmt.Println("New slice: ", newslice)
}
