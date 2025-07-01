package main

import "fmt"

func main() {

	var p *int

	if p == nil {

		fmt.Println("Null pointer:", p)
	}

	var a =457

	var ptr=&a
	
	fmt.Println("value before",*ptr)

	*ptr=500

	fmt.Println("Value after ",*ptr)
}
