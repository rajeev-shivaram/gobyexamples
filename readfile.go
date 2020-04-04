package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	// var file *os.File
	// var err *os.PathError
	file, err := os.Open("daa.txt")
	fmt.Println(reflect.TypeOf(file))
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
