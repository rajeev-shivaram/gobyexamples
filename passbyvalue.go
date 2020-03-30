// In go when you pass values to a function, a copy is passed
// but if we want to modify existing value using functions, a good way
// to do so is by passing those values into functions using.. Pointers!!

package main

import (
	"fmt"
)

func getDoubles(i int) int {
	return 2 * i
}

func getDoublePointers(i *int) int {
	return 2 * *i
}

func main() {
	i := 5
	fmt.Println("before non passbyvalue i.e pass by reference")
	fmt.Println("New values is", getDoubles(i))
	fmt.Println("i value in scope is", i)

	fmt.Println("Pass by reference using pointers is", getDoublePointers(&i))
}
