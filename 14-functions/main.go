package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("GoLang Functions!")

	// adder funtion call
	a := rand.Intn(9)
	b := rand.Intn(9)
	fmt.Println(a, "+", b, "=", adder(a, b))

	// addition of multiple numbers
	fmt.Println(sum(12, 45, 23, 67, 89, 90))
}

// input two values & return result as int
func adder(a int, b int) int {
	return a + b
}

// convert numbers into slice & do addition inside for loop
func sum(numbers ...int) int {
	i := 0
	for _, j := range numbers {
		i += j
	}
	return i
}
