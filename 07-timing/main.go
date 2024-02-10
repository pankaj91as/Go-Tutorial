package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timing function test!")

	// print current date time
	today := time.Now()
	fmt.Printf("Current time is: %v\n", today)

	// format date time
	todayFormated := today.Format("Mon 2 Jan 2006 15:04:05")
	fmt.Printf("Formated Timestamp: %v\n", todayFormated)
}
