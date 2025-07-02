package main

import "fmt"

type name struct{
	string
	int
}

func main() {

	value:= name{
		"aklile",
		25}
	fmt.Println("V: ",value)
	// student1 := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "sklkas",
	// 	age:  24,
	// }
	// fmt.Println("S:", student1)
}
