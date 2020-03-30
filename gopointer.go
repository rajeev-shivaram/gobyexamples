package main

import (
	"fmt"
	"reflect"
)

func createPointer(f *float64) {
	fmt.Print(&f)
}

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var1 := 56.36
	createPointer(&var1)
}
