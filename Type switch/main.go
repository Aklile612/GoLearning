package main

import "fmt"

func main() {

	var a interface{} = "aklile is manutd"

	switch value := a.(type) {
	case int64:
		fmt.Println("This is an int",value)
	case float64:
		fmt.Println("This is a floating number",value)
	case string:
		fmt.Println("This is a string ",value)
	default:
		fmt.Println("Type Unkown",value)
	}
}
