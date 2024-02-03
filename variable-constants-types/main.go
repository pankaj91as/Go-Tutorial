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

	// Variable Declaration Type Five
	myVariableTypeFive, myVariableTypeFiveOne, myVariableTypeFiveTwo := "a", 2, 3.456
	fmt.Println(myVariableTypeFive, myVariableTypeFiveOne, myVariableTypeFiveTwo)
}

func constantsDecleration() {
	// Constant Decleration (You can declare constants with or without data type same as variable)
	const myFirstConstant string = "Golang Coding Challange"
	fmt.Println(myFirstConstant)

	// Complex Data type
	const (
		Red int = iota
		Orange
		Yellow
		Green
		Blue
		Indigo
		Violet
	)
	fmt.Printf("Constant Red has an <%T> data type and value is {%v}\n", Red, Red)
	fmt.Printf("Constant Orange has an <%T> data type and value is {%v}\n", Orange, Orange)
	fmt.Printf("Constant Yellow has an <%T> data type and value is {%v}\n", Yellow, Yellow)
	fmt.Printf("Constant Green has an <%T> data type and value is {%v}\n", Green, Green)
	fmt.Printf("Constant Blue has an <%T> data type and value is {%v}\n", Blue, Blue)
	fmt.Printf("Constant Indigo has an <%T> data type and value is {%v}\n", Indigo, Indigo)
	fmt.Printf("Constant Violet has an <%T> data type and value is {%v}\n", Violet, Violet)
}

func typesDecleration() {
	// String Data type with value
	var stringDataType string = "String Data"
	fmt.Printf("Variable stringDataType has an <%T> data type and value is {%v}\n", stringDataType, stringDataType)

	// Integer Data type
	var integerDataType int = 100
	fmt.Printf("Variable integerDataType has an <%T> data type and value is {%v}\n", integerDataType, integerDataType)

	// Float Data type
	// As per industry standerds do not use numbers in your variable
	// Excuse me here, just to explain clearly I am adding numbers into variables
	var float32DataType float32 = 100.728364982739
	fmt.Printf("Variable float32DataType has an <%T> data type and value is {%v}\n", float32DataType, float32DataType)

	// Uint Data Type
	// Uint is design to take only positive numbers > 0
	var uintDataType uint = 723687462389
	fmt.Printf("Variable uintDataType has an <%T> data type and value is {%v}\n", uintDataType, uintDataType)

	// Boolean Data type
	var booleanDataType bool = true
	fmt.Printf("Variable booleanDataType has an <%T> data type and value is {%v}\n", booleanDataType, booleanDataType)

	// Rune Data type
	var runeDataType rune = 78263482
	fmt.Printf("Variable runeDataType has an <%T> data type and value is {%v}\n", runeDataType, runeDataType)

	// Complex Data type
	var complexDataType complex64 = 12i
	fmt.Printf("Variable complexDataType has an <%T> data type and value is {%v}\n", complexDataType, complexDataType)

	// Complex Data type
	var anyDataType any
	fmt.Printf("Variable anyDataType has an <%T> data type and value is {%v}\n", anyDataType, anyDataType)

	// Complex Data type
	var byteDataType byte
	fmt.Printf("Variable byteDataType has an <%T> data type and value is {%v}\n", byteDataType, byteDataType)

	// Complex Data type
	var errorDataType error
	fmt.Printf("Variable errorDataType has an <%T> data type and value is {%v}\n", errorDataType, errorDataType)

	// Complex Data type
	var uintptrDataType uintptr
	fmt.Printf("Variable uintptrDataType has an <%T> data type and value is {%v}\n", uintptrDataType, uintptrDataType)
}
