package main

import (
	"fmt"
	"log"
	"os"
)

func statCheck() {
	fmt.Println("StatCheck")
	fileInfo, err := os.Stat("hello.go")
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	fmt.Print(fileSize)
}

// func main() {
// 	statCheck()
// }
