// Get the negation of a bool value

package main

import ("fmt")

func negate(b *bool) bool {
	return !(*b)
}

func main(){
	b := true
	fmt.Println("Negative of b is", negate(&b))
}