package main

import "fmt"

// Author struct
type Authour struct {
	name   string
	branch string
	year   int
}

type HR struct {
	details Authour
}

func main() {

	result := HR{
		details: Authour{"aklile", "lafto", 2004},
	}

	fmt.Println("Result:", result)
}
