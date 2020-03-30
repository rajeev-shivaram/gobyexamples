package main

import (
	"fmt"
)

//pointers inside functions
func returnPointer() *float64 {
	flotingNumber := 98.5
	return &flotingNumber
}

func printValAtPointer(poi *float64) {
	fmt.Println(*poi)
}

// func main() {
// 	poi := returnPointer()
// 	fmt.Println(poi)
// 	// prints 0xc000012090

// 	printValAtPointer(poi)
// 	//prints 98.5

// }
