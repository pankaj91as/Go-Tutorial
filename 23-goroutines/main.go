package main

import (
	"fmt"
	"sync"
)

var startNumber int = 0
var wg sync.WaitGroup

func main() {
	// sample goroutines
	go greeter("Hello World!")
	greeter("Welcome to GoLang Tutorials!")

	// goroutines with wait group
	// this example will not give you the expected calculation because
	// go routines are never been in sequense; so addition you can see random numbers
	// to resolve this issue please check next exercise
	i := 0
	for i < 10 {
		wg.Add(1)
		go keepAdding(i)
		i++
	}

	wg.Wait()
}

func greeter(message string) {
	fmt.Println(message)
}

func keepAdding(n int) {
	defer wg.Done()
	startNumber += n
	fmt.Println(startNumber)
}
