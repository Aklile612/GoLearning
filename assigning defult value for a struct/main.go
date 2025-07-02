package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) Instance() {
	if s.name == "" {
		s.name = "Robert"
	}

	if s.age == 0 {
		s.age = 25
	}
}

func main() {
	student1 := student{}

	student1.Instance()
	fmt.Println(student1)

	student2:= student{
		name: "aklile",
	}
	student2.Instance()
	fmt.Println(student2)
}
