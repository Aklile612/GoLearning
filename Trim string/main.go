package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := "!! welcome buddy aagain and lets go"

	trim1 := strings.Trim(result1, "!")

	fmt.Println(trim1)

	//trim space

	result2:= "   welcome to the hustle   "

	trim2:= strings.TrimSpace(result2)

	fmt.Println(trim2)

	//trim suffix

}
