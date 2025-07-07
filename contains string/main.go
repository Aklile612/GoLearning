package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Abebe", "be"))

	fmt.Println(strings.Contains("Germany","GER"))

	//Fields

	message:="Abebe Beso Libela Blo"

	messageArray:=strings.Fields(message)

	fmt.Println(messageArray)
}
