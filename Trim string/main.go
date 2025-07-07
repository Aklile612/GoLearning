package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := "!! welcome buddy aagain and lets go"

	trim1 := strings.Trim(result1, "!")

	fmt.Println(trim1)
}
