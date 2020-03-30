package main

import (
	"fmt"
)

func paintNeeded1(width, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height <= 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	return (height * width) / 10.0, nil
}

// func main() {
// 	paints, err := paintNeeded1(0.0, 11.256)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%.2f liters needed.", paints)
// }
