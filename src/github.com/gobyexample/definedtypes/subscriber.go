package main

import "fmt"

type subscriber struct {
	name string
	rate float64
	active bool
}

func applyDiscount(s *subscriber){
	s.rate = 4.49
}

func main(){
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)

	fmt.Println(s)
}