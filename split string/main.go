package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "This,is,the,fact,that,we,are,tired"

	str2 := "THis should be the only way"

	result1 := strings.Split(str1, ",")
	result2 := strings.Split(str2, " ")
	fmt.Println(result1)
	fmt.Println(result2)
}