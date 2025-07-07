package main

import "fmt"
func main() {
	var name string
	var age int
	var height float64
	var isAbsent bool

	fmt.Scan(&name)
	fmt.Scan(&height)
	fmt.Scan(&age)
	fmt.Scan(&isAbsent)

	fmt.Printf("%s %d %f %t", name,age,height,isAbsent)
}