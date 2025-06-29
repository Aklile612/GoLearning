package main

import "fmt"

func adds(x int, y int) int {

	result := x + y

	return result
}

func main() {

	sum := adds(20, 30)

	fmt.Printf("result : %d\n", sum)
}
