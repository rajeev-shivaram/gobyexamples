package main

import (
	"fmt"
	"math"
)

func maximum(nums ...float64) {
	
	fmt.Println("Numbers are", nums)
	maxi := math.Inf(-1)

	for _, num := range nums {
		if num > maxi {
			maxi = num
		}
	}

	fmt.Println("MAx is", maxi)
}

func main() {
	maximum(5, 4, 3, 2, 1)
	nums := []float64{5, 4, 3, 2, 1}
	maximum(nums...)
}
