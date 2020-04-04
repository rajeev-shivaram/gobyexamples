package main

import (
	"fmt"
	"github.com/gobyexample/prose"
)

func main(){
	phares := []string{"apples", "oranges"}
	fmt.Println("a bowl of", prose.JoinWithCommas(phares))
}