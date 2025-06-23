package main

import "fmt"

func main() {

	a := []int{10, 20, 30}

	b := make([]int, 3, 5)

	copy(b, a)

	fmt.Println("slice b", b)
}
