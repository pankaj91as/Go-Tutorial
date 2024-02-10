package main

import (
	"fmt"
)

type Users struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	fmt.Println("Methods in GoLang!")

	// Add values into struct
	user := Users{"Pankaj", 34, "pankaj91as@gmail.com", true}
	fmt.Println(user)

	// get user email
	fmt.Println(user.GetUserEmail())
}

func (u Users) GetUserEmail() string {
	return u.Email
}
