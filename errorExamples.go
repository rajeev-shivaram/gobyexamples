//Example for an error
package main

import (
	"fmt"
	"log"
	"math"
)

func errorhandler() {
	height := 2.569855236
	err := fmt.Errorf("height %2.5f cant be negative", height)
	// fmt.Println(err.Error()) //Prints the error message
	log.Fatal(err) //prints the error message again and exits the program
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

// func main() {
// 	errorhandler()
// 	manyReturns()
// }
