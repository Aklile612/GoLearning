package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := strings.Compare("Hello", "Hello")

	fmt.Println(result1)

	result2:= strings.Compare("Hello","Aklile")
	fmt.Println(result2)
	result3 := strings.Compare("Hello", "hello")

	fmt.Println(result3)

}
