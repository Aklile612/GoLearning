package main

import "fmt"

func main() {

	empolyee := make(map[string]int)

	empolyee["A"] = 1
	empolyee["B"] = 10
	empolyee["C"] = 19
	empolyee["D"] = 20

	empolyeeList := make(map[string]int)

	fmt.Println("Employee: ", len(empolyee))
	fmt.Println("Value for A is",empolyee["B"])
	fmt.Println("Employee List", len(empolyeeList))

}
