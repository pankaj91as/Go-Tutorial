package main

import "fmt"

func main() {
	fmt.Println("Maps (HashTable) in GoLang!")

	myHashTable := make(map[string]string)
	fmt.Printf("Map Default Value: %v", myHashTable)

	// Insert values into MAP
	myHashTable["01"] = "First"
	myHashTable["02"] = "Second"
	myHashTable["03"] = "Third"
	myHashTable["04"] = "Fourth"
	myHashTable["05"] = "Fifth"

	fmt.Println("Inserted some values into map!")
	fmt.Println(myHashTable)

	// Delete specific key value from map
	delete(myHashTable, "04")
	fmt.Println(myHashTable)
}
