package main

import "fmt"

func main() {

	var slice1 = make([]int, 4, 7)

	fmt.Printf("slice1= %v\t, Lenght= %d\t , capacity=%d\n", slice1, len(slice1), cap(slice1))

	slice2:= make([]int,4)
	fmt.Printf("slice2= %v\t, Lenght= %d\t , capacity=%d\n", slice2,len(slice2),cap(slice2))

}
