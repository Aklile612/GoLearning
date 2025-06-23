package main

import "fmt"

func main() {

	a := make([]int, 2, 5)

	a[0] = 2
	a[1] = 5

	fmt.Println("Slice before", a)

	a=append(a, 3,4,5,6,7,8)

	fmt.Println("Slice after", a)

}
