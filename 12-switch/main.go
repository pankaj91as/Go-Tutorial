package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("GoLang Switch Case!")

	// Generate random number
	diceNumber := rand.Intn(6) + 1

	// Switch case
	switch diceNumber {
	case 1:
		fmt.Println("One!")
		break
	case 2:
		fmt.Println("Two!")
		break
	case 3:
		fmt.Println("Three!")
		fallthrough // fallthrough to next case
	case 4:
		fmt.Println("Four!")
		break
	case 5:
		fmt.Println("Five!")
		break
	case 6:
		fmt.Println("Six!")
		break
	default:
		fmt.Println("Roll your dice again!!!")
	}
}
