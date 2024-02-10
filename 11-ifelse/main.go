package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Golang If Else Condition!")

	// Generate random integer
	var luckyNumber = rand.Intn(100)

	if luckyNumber%2 == 0 { // check if devisible by two
		fmt.Printf("My really luck number is %v because its devisible by 2!\n", luckyNumber)
	} else if luckyNumber%2 == 0 && isPrime(luckyNumber) { // check else if luckyNumber devisible by 2 & it's a prime number or not
		fmt.Printf("My really luck number is %v because its prime number!\n", luckyNumber)
	} else { // else print number
		fmt.Printf("%v Is not my lucky number... try one more time!", luckyNumber)
	}
}

// prime number varification function
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
