package main

import (
	"fmt"

	"github.com/gobyexample/datafile"
)

func main() {
	requiredItemsInKgsWeek := [7]float64{52.2, 12.4, 78.9, 52.1, 12.4, 78.9, 63.2}

	var total float64 = 0
	for _, value := range requiredItemsInKgsWeek {
		total += value
	}
	ff, err := datafile.GetFloats("E:\\daa.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ff)

	fmt.Printf("%5.2f", float64(total)/float64(len(requiredItemsInKgsWeek)))

}
