package main

import "fmt"

func main() {

	a := 10

	b := 10

	if (a < b) {
		fmt.Println("A less than B")
	}else if (a==b){
		fmt.Println("A equal to B")
	}else{
		fmt.Println("A greater than B")
	}
}
