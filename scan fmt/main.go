package main

import "fmt"
func main() {
	var name string
	var age int
	var height float64
	var isAbsent bool

	fmt.Scanln(&name)
	fmt.Scanln(&height)
	fmt.Scanln(&age)
	fmt.Scanln(&isAbsent)

	fmt.Printf("%s %d %f %t", name,age,height,isAbsent)
}