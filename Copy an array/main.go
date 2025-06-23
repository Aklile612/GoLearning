package main

import "fmt"

func main() {

	array1 := [3]string{"a", "b", "c"}

	array2 := array1

	array3 := &array1

	fmt.Printf("Array1: %v\n", array1)
	fmt.Printf("Array1: %v\n", array2)

	array1[0]="z"
	fmt.Printf("Array1: %v\n", array2)
	fmt.Printf("Array1: %v\n", *array3)
}
