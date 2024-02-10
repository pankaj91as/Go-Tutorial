package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web Requests in GoLang!")

	// GET url
	url := "https://pankaj91as.github.io/sample-html/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// type of http response
	fmt.Printf("%T\n", response)

	// print response body
	pageData, _ := io.ReadAll(response.Body)
	fmt.Println(string(pageData))
}
