package datafile

import (
	"log"
	"bufio"
	"os"
	"strconv"
)

// GetFloats doesn
func GetFloats(filename string) ([3]float64, error){
	var numbers [3]float64
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	err = file.Close()
	if (err != nil || scanner.Err() != nil) {
		log.Fatal(err)
	}
	return numbers, nil
}