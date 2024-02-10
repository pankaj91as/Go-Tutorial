package main

import "fmt"

func main() {
	// Create Variable
	var myTestVar = 2312
	fmt.Printf("Variable value: %v\n", myTestVar)

	// create pointer
	var myFirstPointer = &myTestVar
	fmt.Printf("Memory address of variable: %v\n", myFirstPointer)

	// reterive Value from pointer
	var myPointerValue = *myFirstPointer
	fmt.Printf("Reterive value from pointer: %v\n", myPointerValue)

	// another type to define pointer variable
	var defineSecondPointer *int
	fmt.Println(defineSecondPointer)
}
