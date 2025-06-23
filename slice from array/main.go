package main

import "fmt"

func main() {

	array := [5]int{1, 3, 4, 5, 6}

	newslice := array[3:5]

	fmt.Println("new slice", newslice)
}
