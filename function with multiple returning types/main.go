package main

import "fmt"

func rectangle(w int, h int) (int, int) {

	perimeter := 2 * (w + h)

	area := w * h

	return area, perimeter
}

func main() {

	a, p := rectangle(3, 6)

	fmt.Println("Area:", a)

	fmt.Println("Permeter: ", p)
}
