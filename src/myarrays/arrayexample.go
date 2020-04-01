package myarrays

import (
	"fmt"
	"time"
)

// ArrayTimes displays arrays of init and uninit times
func ArrayTimes() {
	var myArray [3]time.Time

	fmt.Println("Before init")
	fmt.Println(myArray)
	fmt.Println("After init")

	myArray[0] = time.Unix(1257894000, 0)

	fmt.Println(myArray)
}


// Prints the following
// Before init
// [0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC]
// After init
// [2009-11-11 04:30:00 +0530 IST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC]

