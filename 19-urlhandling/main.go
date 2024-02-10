package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL parsing GoLang!")

	// sample URL
	myUrl := "https://pankaj91as.github.io:443/sample-html/?v=7826348293"
	parsedURL, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	// get url scheme
	fmt.Println("Scheme: ", parsedURL.Scheme)

	// get host from URL
	fmt.Println("Host: ", parsedURL.Host)

	// fetch port from url
	fmt.Println("Port: ", parsedURL.Port())

	// fetch path params
	fmt.Println("Path param: ", parsedURL.Path)

	// get url query params
	queryParams := parsedURL.Query()
	for _, value := range queryParams {
		fmt.Println("Query Parameters: ", value)
	}

	// fetch specific value from query params
	fmt.Printf("Value of query parameter v is: %v\n", queryParams["v"])
}
