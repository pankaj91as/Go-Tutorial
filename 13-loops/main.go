package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang!")

	// for loop simple
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// for loop range
	var intRange = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, j := range intRange {
		fmt.Println(j)
	}

	// foreach type loop
	rangeValue := 0
	for rangeValue < 10 {

		if rangeValue == 4 {
			rangeValue++
			continue
		}

		// goto statement
		if rangeValue == 9 {
			goto birthday
		}

		fmt.Printf("Value is: %d\n", rangeValue)
		rangeValue++
	}

birthday:
	fmt.Println("Its my birthdate!")
}
