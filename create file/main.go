package main

import (
	"log"
	"os"
)

func main() {
	emptyFile,err:= os.Create("empty.txt")

	if err!= nil{
		log.Fatal(err)
	}else{
		log.Println(emptyFile)
	}
}