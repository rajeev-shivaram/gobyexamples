package main

import (
	"github.com/gobyexample/datafile"
	"fmt"
	"log"
)

func main(){
	lines, err := datafile.GetString("E:\\data.txt")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(lines)
}