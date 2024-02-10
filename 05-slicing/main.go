package main

import "fmt"

func main() {
	// Define slice variable
	var cars = []string{"Audi", "BMW", "Mercedes"}
	fmt.Println("Slice variable cars having: ", cars)

	// Define slice from array variable
	var carsArray = [5]string{"Audi", "BMW", "Mercedes", "Maruti", "TATA"}
	carsSlice := carsArray[3:5]
	fmt.Printf("My Array is: %v\n", carsArray)
	fmt.Printf("New Slice made from array is: %v\n", carsSlice)

	// Slice With Length & Capacity
	makeSliceWithLengthAndCapacity := make([]int, 5, 100)
	fmt.Printf("makeSliceWithLengthAndCapacity = %v\n", makeSliceWithLengthAndCapacity)
	fmt.Printf("length = %d\n", len(makeSliceWithLengthAndCapacity))
	fmt.Printf("capacity = %d\n", cap(makeSliceWithLengthAndCapacity))

	// with omitted capacity
	makeSliceWithLengthOnly := make([]int, 5)
	fmt.Printf("makeSliceWithLengthOnly = %v\n", makeSliceWithLengthOnly)
	fmt.Printf("length = %d\n", len(makeSliceWithLengthOnly))
	fmt.Printf("capacity = %d\n", cap(makeSliceWithLengthOnly))
}
