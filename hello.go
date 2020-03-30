package main

import (
	"fmt"
	"strings"
)

func hello() {
	// price := 100
	// fmt.Println("Price is", price, "dollars")

	// var taxRate float64 = 0.08
	// var tax float64 = float64(price) * taxRate
	// fmt.Println("Tax is", tax, "dollars")

	// var total float64 = float64(price) + tax
	// fmt.Println("Total is", total, "dollars")

	// var availableFunds int = 120
	// fmt.Println(availableFunds, "dollars available.")
	// fmt.Println("Within budget?", float64(availableFunds) >= total)
	// ---------------------------------------------------------------
	// var now time.Time = time.Now()

	// var hour, minute, sec int
	// hour, minute, sec = now.Clock()
	// fmt.Println("H:", hour, "\nminute:", minute, "\nsec", sec)
	// fmt.Println(reflect.TypeOf(now.Year()))

	// //LocalTime := now.Local()
	// var year, todaysDate int
	// var thismonth time.Month
	// year, thismonth, todaysDate = now.Date()
	// fmt.Println("year:", year, "\nMonth: ", thismonth, "\nDate", todaysDate)

	//Replacer coder
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	fmt.Println("OK!")

}
