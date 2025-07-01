package main

import "fmt"

func main() {
	a := 20

	var p *int

	p = &a

	fmt.Println(&a)

	fmt.Println(p)

	fmt.Println(*p)

}
