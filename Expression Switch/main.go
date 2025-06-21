package main

import "fmt"

func main() {
	a := 20

	switch {
	case a == 1:
		fmt.Println("The first day")
	case a == 2:
		fmt.Println("the second day")
	default:
		fmt.Println("This is the End hold your breath that we make ")
	}
}
