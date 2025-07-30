package main

import "fmt"

func main() {

	var x int = 0

lable1:
	for x < 8 {

		if x == 5 {
			x++
			goto lable1
		}
		fmt.Println(x,"")
		x++
	}
}
