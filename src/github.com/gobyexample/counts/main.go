package main

import (
	"fmt"
	"log"

	"github.com/gobyexample/datafile"
)

func main() {
	lines, err := datafile.GetString("E:\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// counts := make(map[string]int)
	for _, line := range(lines) {
		fmt.Println(line)
	}
	// for key, value := range(counts) {
	// 	fmt.Println("Name:",key,"value:",value)
	// }
}
