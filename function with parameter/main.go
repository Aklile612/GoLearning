package main

import "fmt"

func add(x int, y int) {

	result := x + y

	fmt.Printf("the value of the sum is: %d \n", result)
}

func main(){
	add(12,13)
}