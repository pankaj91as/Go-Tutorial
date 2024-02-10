package main

import (
	"fmt"
)

type UserStruct struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

type UserStructWithArray struct {
	Name   string
	Age    int
	Email  string
	Status bool
	Trans  []int
}

type Store struct {
	Data map[int]int
}

func main() {
	fmt.Println("Structs in GoLang!")

	// Blank struct default values
	fmt.Println("Default value of blank struct")
	fmt.Println(UserStruct{})

	// Add values into struct
	user := UserStruct{"Pankaj", 34, "pankaj91as@gmail.com", true}
	fmt.Println(user)

	// print with structure
	fmt.Printf("Details of users are: %+v\n", user)

	// print specific value
	fmt.Printf("User Name is: %v\n", user.Name)

	// Add values into struct with array
	userWithArray := UserStructWithArray{Name: "Pankaj", Age: 34, Email: "pankaj91as@gmail.com", Status: true, Trans: []int{1}}
	fmt.Println(userWithArray)

	// Add values into map struct
	var i = 0
	s := Store{
		Data: make(map[int]int),
	}
	for i = 0; i < 10; i++ {
		s.Data[i] = i
	}

	// Print map struct
	fmt.Println(s)

	// Fetch specific/index value from map struct
	fmt.Println(s.Data[5])
}
