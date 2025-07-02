package main

import "fmt"

type address struct {
	city     string
	location string
	zipCode  int
}

func main() {
	a1:= address{"addis","tofla",2033}
	
	fmt.Println(a1)

	a2:= address{city: "california",location: "paradise",zipCode: 2000}

	fmt.Println("Address2",a2)

	a3:=new(address)

	a3.city="egyipr"
	a3.location="location"
	a3.zipCode=5059

	fmt.Println(a3)
}
