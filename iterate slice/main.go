package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40}

	for index, value := range slice {
		fmt.Println("index,value", index, value)
	}
}
