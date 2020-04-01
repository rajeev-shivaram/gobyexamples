package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	var total float64 = 0
	for _, argument := range args {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		total += number
	}

	fmt.Printf("%5.2f", float64(total)/float64(len(args)))

}
