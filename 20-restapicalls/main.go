package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GoLang GET Request Example!")

	// GET user url
	SampleURL := "https://reqres.in/api/users/5"

	// Make HTTP GET call
	response, err := http.Get(SampleURL)
	if err != nil {
		panic(err)
	}

	// print direct response
	// fmt.Println(response)

	// Read Byte Data
	bodyData, err := io.ReadAll(response.Body)
	fmt.Println(string(bodyData))

	// another method
	var responseString strings.Builder
	byteCount, err := responseString.Write(bodyData)
	fmt.Printf("No of bytes data %d\n", byteCount)
	fmt.Println(responseString.String())

	// GoLang Post Request
	fmt.Println("GoLang POST Request Example!")

	// POST Endpoint
	postURL := "https://reqres.in/api/register"

	// POST Body Data
	requestBody := strings.NewReader(`{
		"email": "eve.holt@reqres.in",
		"password": "pistol"
	}`)

	// Make POST Request
	request, err := http.NewRequest("POST", postURL, requestBody)
	if err != nil {
		panic(err)
	}
	request.Header.Set("content-type", "application/json; charset=utf-8")
	request.Header.Set("access-control-allow-origin", "*")
	responseNew, err := http.DefaultClient.Do(request)
	postResponse, err := io.ReadAll(responseNew.Body)
	fmt.Println(string(postResponse))

	// GoLang PUT Request
	fmt.Println("GoLang PUT Request Example!")

	// PUT Endpoint
	putURL := "https://reqres.in/api/users/2"

	// POST Body Data
	putRequestBody := strings.NewReader(`{
		"name": "morpheus",
		"job": "zion resident"
	}`)

	// Make POST Request
	putRequest, err := http.NewRequest("PUT", putURL, putRequestBody)
	if err != nil {
		panic(err)
	}
	putRequest.Header.Set("content-type", "application/json; charset=utf-8")
	putRequest.Header.Set("access-control-allow-origin", "*")
	putResponseNew, err := http.DefaultClient.Do(putRequest)
	putResponse, err := io.ReadAll(putResponseNew.Body)
	fmt.Println(string(putResponse))

	// GoLang PATCH Request
	fmt.Println("GoLang patch Request Example!")

	// patch Endpoint
	patchURL := "https://reqres.in/api/users/2"

	// POST Body Data
	patchRequestBody := strings.NewReader(`{
		"name": "morpheus",
		"job": "zion resident"
	}`)

	// Make POST Request
	patchRequest, err := http.NewRequest("PATCH", patchURL, patchRequestBody)
	if err != nil {
		panic(err)
	}
	patchRequest.Header.Set("content-type", "application/json; charset=utf-8")
	patchRequest.Header.Set("access-control-allow-origin", "*")
	patchResponseNew, err := http.DefaultClient.Do(patchRequest)
	patchResponse, err := io.ReadAll(patchResponseNew.Body)
	fmt.Println(string(patchResponse))
}
