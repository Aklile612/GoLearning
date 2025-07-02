package main

import "fmt"

type work struct {
	branch string
	name   string
	salary int
}

func (a work) show() {

	fmt.Println("Branch: ", a.branch)
}
func main() {
	result := work{
		branch: "lafto",
		name:   "aklile",
		salary: 2000,
	}

	result.show()
}
