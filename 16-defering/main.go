package main

import "fmt"

func main() {
	fmt.Println("Defer Functionality in GoLang!")

	// defer LIFO
	fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Third")
	defer fmt.Println("Fourth")

}
