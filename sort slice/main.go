package main

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := []int{10, 42, 2, 4, 43, 60}

	slice2 := []string{"java", "go", "c++", "react"}

	sort.Ints(slice1)
	sort.Strings(slice2)
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("slice2: ", slice2)
}
