package main

import "fmt"

func main() {
	fmt.Println("GoLang Panic & Recover!!!")
	functionOne()
	fmt.Println("panic recovered!")
}

func functionOne() {
	fmt.Println("Call my controller...")
	// defer func() {
	// 	if rec := recover(); rec != nil {
	// 		fmt.Println("recover faild in functionTwo :(")
	// 	}
	// }()
	functionTwo()
	fmt.Println("function one executed successfuly!")
}

func functionTwo() {
	fmt.Println("Retriving data from storage....")
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("recover faild in functionTwo :(")
		}
	}()
	if true {
		panic(fmt.Errorf("Panic I got in function %d", 2))
	}
	fmt.Println("Returning data from functionTwo")
}
