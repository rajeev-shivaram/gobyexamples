// pass_fail report whether a grade is passing or failing

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passFail() {
	fmt.Print("Pass fail go grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	parsedInput := strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(parsedInput, 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Print("Status:", status)
}

// func main() {
// 	passFail()
// }
