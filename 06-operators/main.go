package main

import "fmt"

func main() {
	var a int = 5
	var b int = 10

	var c string = "Hello"
	var d string = "World!"

	// Artithmatic Operators
	fmt.Println(a + b) // Addition (+) sign is used for addition
	fmt.Println(a - b) // Subtraction (-) sign is used for subtraction
	fmt.Println(a * b) // * sign is used for multiplication
	fmt.Println(a / b) // / sign is used for division
	fmt.Println(a % b) // % sign is used for division reminder
	a++
	fmt.Println(a) // Increment (++) is used to Increment number by one
	a--
	fmt.Println(a) // Decrement (--) is used to Decrement number by one

	// Assignment Operators
	var e int = 5

	var f = 0
	f += 14

	var g = 5
	g -= 23

	var h = 5
	h /= 32

	var i = 5
	i %= 41

	var j = 5
	j *= 50

	var k = 5
	k &= 5

	var l = 5
	l |= 10

	var m = 5
	m ^= 15

	var n = 5
	n >>= 20

	var o = 5
	o <<= 25

	fmt.Print("Assignment Operators: \n")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(0)

	// Comparison Operators
	fmt.Println(a == b)      // == sign is used for comparison
	fmt.Println(a != b)      // == sign is used for comparison
	fmt.Println(a > b)       // == sign is used for comparison
	fmt.Println(a < b)       // == sign is used for comparison
	fmt.Println(a >= b)      // == sign is used for comparison
	fmt.Println(a <= b)      // == sign is used for comparison
	fmt.Println(c + " " + d) // + sign is used for concatination
}
