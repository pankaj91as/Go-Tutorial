// Please observ the input given to variable & output showing in console is very important
// to understand the diffrence & functionality
package main

import "fmt"

// Global Variable Decleration (Data type is optional)
var myGlobalVariable string = "JWToken"

func main() {
	// Call Variable Declaration Function
	fmt.Println("Variable Decleration Output Below:")
	variableDecleration()
	fmt.Print("\n\n\n")

	// Call Constant Declaration Function
	fmt.Println("Constants Decleration Output Below:")
	constantsDecleration()
	fmt.Print("\n\n\n")

	// Type Decleration
	fmt.Println("Data Type Decleration Output Below:")
	typesDecleration()
	fmt.Print("\n\n\n")
}

func variableDecleration() {
	// Variable declaration type 1 (With Data Type)
	var myVariableTypeOne string = "Variable Type One"
	fmt.Println(myVariableTypeOne)

	// Variable declaration type 2 (Without Data Type)
	var myVariableTypeTwo = "Variable Type Two"
	fmt.Println(myVariableTypeTwo)

	// Variable declaration type 3 (Auto predict data type from value)
	myVariableTypeThree := "Variable Type Three"
	fmt.Println(myVariableTypeThree)

	// Variable Declaration type 4 (Global)
	fmt.Println(myGlobalVariable)
}

func constantsDecleration() {
	// Constant Decleration (You can declare constants with or without data type same as variable)
	const myFirstConstant string = "Golang Coding Challange"
	fmt.Println(myFirstConstant)
}

func typesDecleration() {
	// String Data type with value
	var stringDataType string = "String Data"
	fmt.Println(stringDataType)

	// String Data type without value
	var stringDataTypeWithoutValue string
	fmt.Println(stringDataTypeWithoutValue)

	// Integer Data type
	var integerDataType int = 100
	fmt.Println(integerDataType)

	// Integer Data type without value
	var integerDataTypeWithoutValue int
	fmt.Println(integerDataTypeWithoutValue)

	// Float Data type
	// As per industry standerds do not use numbers in your variable
	// Excuse me here, just to explain clearly I am adding numbers into variables
	var float32DataType float32 = 100.728364982739
	fmt.Println(float32DataType)

	// Float Data type without value
	var float32DataTypeWithoutValue float32
	fmt.Println(float32DataTypeWithoutValue)
}
