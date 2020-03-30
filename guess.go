// Random number Guesser Game

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guesser() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	target := rand.Intn(10) + 1
	fmt.Println("I've choosen a random number between 1 and 100")
	fmt.Println("Can you guess it ?")

	// fmt.Println(target)

	success := false
	for guess := 0; guess < 10; guess++ {

		fmt.Println("You have", 10-guess, "guesses left.")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make a guess:")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input := strings.TrimSpace(userInput)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//Compare the guess and target
		if target > guess {
			fmt.Println("Oops. Your guess was LOW.")
		} else if target < guess {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Yay. You got it RIGHT!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("You didnt guess the correct number, it was", target)
	}

}

// func main() {
// 	guesser()
// }
