package main

import "fmt"

func main() {
	// printing functions
	string := "string"
	fmt.Println("This string prints with fmt.Println function!") // function fmt.Println adding \n (new line) at the end of line
	fmt.Print("This string prints with fmt.Print function!")
	fmt.Printf("This %T prints with %v function!", string, "fmt.Printf")

	number := 15.7
	text := "Hello World!"

	// number formating verbs
	fmt.Println()
	fmt.Printf("%v\n", number)   // print value
	fmt.Printf("%#v\n", number)  // print value in golang format
	fmt.Printf("%T\n", number)   // print type of variable
	fmt.Printf("%v%%\n", number) // print percent sign with number

	// string formating verbs
	fmt.Printf("%v\n", text)   // print value
	fmt.Printf("%#v\n", text)  // print value in golang format
	fmt.Printf("%T\n", text)   // print type of variable
	fmt.Printf("%v%%\n", text) // print percent sign with number

	// Integer formating verbs
	integer := 20
	binaryInteger := 10110
	fmt.Printf("%b\n", integer)       // convert number into binary format
	fmt.Printf("%d\n", binaryInteger) // print decimal format
	fmt.Printf("%+d\n", integer)      // add plus (+) sign before number
	fmt.Printf("%o\n", integer)       // convert to octal number
	fmt.Printf("%O\n", integer)       // convert to octal number and adding Oo in front of number
	fmt.Printf("%x\n", integer)       // convert
	fmt.Printf("%X\n", integer)       // Prints the value as hex dump of byte values
	fmt.Printf("%#x\n", integer)      // Prints the value as hex dump with spaces
	fmt.Printf("%4d\n", integer)      // Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", integer)     // Pad with spaces (width 4, left justified)
	fmt.Printf("%05d\n", integer)     // making x number of integer by prepending 0

	// String formating verbs
	mystring := "Golang Tutorial"
	fmt.Printf("%s\n", mystring)   // Prints the value as plain string
	fmt.Printf("%q\n", mystring)   // Prints the value as a double-quoted string
	fmt.Printf("%8s\n", mystring)  // Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", mystring) // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", mystring)   // Prints the value as hex dump of byte values
	fmt.Printf("% x\n", mystring)  // Prints the value as hex dump with spaces

	// Boolean formating verbs
	myTrueBool := true
	myFalseBool := false
	fmt.Printf("%t\n", myTrueBool)  // Value of the boolean operator in true or false format (same as using %v)
	fmt.Printf("%t\n", myFalseBool) // Value of the boolean operator in true or false format (same as using %v)

	// Float formating verbs
	myFloat := 14.29
	fmt.Printf("%e\n", myFloat)     // Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", myFloat)     // Decimal point, no exponent
	fmt.Printf("%.2f\n", myFloat)   // Default width, precision 2
	fmt.Printf("%.9f\n", myFloat)   // Default width, precision 9
	fmt.Printf("%15.2f\n", myFloat) // Width 15, precision 2
	fmt.Printf("%g\n", myFloat)     // Exponent as needed, only necessary digits
}
