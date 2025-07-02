package main

import "fmt"

func main() {
	var result1 =" Go shut ur damn mouth"
	display(result1)
	
	result2:=1223
	display(result2)
}
func display(a interface{}) {
	value,ok := a.(string)

	fmt.Printf("Value:%v --- status is %v\n ", value,ok)
}
