package main

import "fmt"

func main() {

	slice1 := []int{10, 20, 30,90,100}

	slice2 := []int{40, 50, 60}

	slice2 = append(slice2, slice1...)

	fmt.Println("slice 2", slice2)
}
