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

	fmt.Println("Employee List", len(empolyeeList))

	empolyee["E"] = 30

	empolyee["F"] = 40

	fmt.Println("new", empolyee)

	delete(empolyee, "F")

	fmt.Println("after", empolyee)

	for index, value := range empolyee {

		fmt.Printf("index:%s\t Value:%d\n", index, value)
	}

}
