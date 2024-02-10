package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("GoLang File RW Functionality!")

	// Create file using os.Create function
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}

	// Close file after all operation finished
	defer file.Close()

	// output of file creation
	fmt.Println(file)

	// Write Data Into File using io.WriteString function
	myString := "I love GoLang!!!"
	number, err := io.WriteString(file, myString)
	fmt.Println(number)

	// Read File Using os.ReadFile package
	data, err := os.ReadFile("./myfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//
	// Read File Using os.Open package
	fileRead, err := os.Open("./myfile.txt")
	if err != nil {
		panic(err)
	}

	ioData, err := io.ReadAll(fileRead)
	if err != nil {
		log.Fatal(err)
	}

	// Print the file contents.
	fmt.Println(string(ioData))

}
