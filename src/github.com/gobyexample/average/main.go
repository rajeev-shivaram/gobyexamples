package main

import "fmt"

func main() {
	requiredItemsInKgsWeek := [7]float64{52.2, 12.4, 78.9, 52.1, 12.4, 78.9, 63.2}

	var total float64 = 0
	for _, value := range requiredItemsInKgsWeek {
		total += value
	}

	fmt.Printf("%5.2f", total)

}
