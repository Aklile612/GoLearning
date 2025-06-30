package main

import "fmt"

func main() {

	first := map[int]string{

		1: "a",
		2: "b",
		3: "c",
	}
	second := map[int]string{
		4: "d",
		5: "e",
		6: "f",
	}
	for i, j := range second {
		first[i] = j
	}
	fmt.Println(first)
}
