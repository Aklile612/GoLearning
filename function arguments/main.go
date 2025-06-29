package main

import "fmt"

func swap(x, y int) int {

	temp := x

	x = y
	y = temp

	return temp
}

func main() {
	var a int = 100

	var b int = 200

	fmt.Println("value of A: ", a)

	fmt.Println("value of B: ", b)

	swap(a, b)

	fmt.Println("value of A: ", a)

	fmt.Println("value of B: ", b)

}
