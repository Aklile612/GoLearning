package main

import "fmt"

func main() {
	student1 := struct {
		name string
		age  int
	}{
		name: "sklkas",
		age:  24,
	}
	fmt.Println("S:", student1)
}
