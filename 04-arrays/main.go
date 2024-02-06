package main

import "fmt"

func main() {
	// String array decleration
	var cars = [3]string{"Audi", "BMW", "Mercedes"}
	fmt.Println("Arrays Of Car Name: ", cars)
	fmt.Printf("Type of cars varible is %T\n", cars)

	// Integer array decleration with inferred lengths
	var marks = [...]int{50, 40, 60, 30, 90, 98}
	fmt.Println("Arrays Of marks: ", marks)
	fmt.Printf("Type of marks varible is %T\n", marks)

	// Access specific value form array using index
	fmt.Println("Car name on first [1] position is: ", cars[0])
}
