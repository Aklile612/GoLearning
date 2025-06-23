package main

import "fmt"

func main() {

	var i int = 0

	for i < 8 {

		fmt.Printf("variable= %d \n", i)

		if (i==5){
			i=i+2

			continue
		}
		i++
	}
}
