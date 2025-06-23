package main

import "fmt"

func main() {

	var students [3]string

	students[0] = "aklile"

	students[1] = "abebe"

	students[2] = "bekele"

	for i := 0; i < len(students); i++ {
		fmt.Println("names: ", students[i])
	}


	players:= [3]string{"Ronaldo","suiiiii","goat"}

	for i:=0;i<3;i++{
		fmt.Println(players[i])
	}
}
