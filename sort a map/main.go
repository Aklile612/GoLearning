package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedArray := make(map[int]string)

	unsortedArray[50] = "A"
	unsortedArray[30] = "C"
	unsortedArray[40] = "B"

	unsortedArray[20] = "D"

	arr := make([]int, 0, len(unsortedArray))

	for i := range unsortedArray {
		arr = append(arr, i)
	}

	sort.Ints(arr)

	for _, values := range arr {
		fmt.Println(unsortedArray[values], values)
	}

	stringArr:=make([]string,0,len(unsortedArray))

	for _,values := range unsortedArray{
		stringArr = append(stringArr, values)
		
	}

	sort.Strings(stringArr)

	for _,values := range stringArr{
		fmt.Println(values)
	}
}
