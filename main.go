package main

import "fmt"

// type Liters float64
// type Gallons float64

type myType string

func (m myType) sayHi() {
	fmt.Println("Hi from", m)
}


func main() {
	// var carFuel Gallons
	// var busFuel Liters

	// carFuel = Gallons(10.0) //convert float64 to Gallons type
	// busFuel = Liters(240.0) //convert float64 to Liters type

	// fmt.Println(carFuel, busFuel)
	// value := myType("a value")
	// value.sayHi()

}