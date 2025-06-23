package main


import "fmt"

func main() {

	names := [...]int64{10, 20, 40}

	fmt.Println("Lenght", len(names))

	x:= [5] int {0:10,2:30}

	fmt.Println(x)

	y:= [5] string{"ak","li","le","bo","ss"}

	fmt.Println(y[:3])
}
