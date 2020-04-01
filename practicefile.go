package main

import "fmt"

func main() {
	myarray := [3]int{1, 2, 3}
	fmt.Println(myarray)

	text := []string{
		"hello",
		"world",
		"whats",
		"up",
	}

	for i:=0; i<len(text);i++ {
		fmt.Println(text[i])
	}

	for index, value := range(text) {
		fmt.Println(index, value)
	}

}
