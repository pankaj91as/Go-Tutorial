package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Create Context
	// context timeout set to low to verify dedline
	// increase timeout & verify the output
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	// create http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://dummyimage.com/1900x800/000000/fff.png", nil)
	if err != nil {
		panic(err)
	}

	// Perform HTTP Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// verify the response of request
	data, err := io.ReadAll(res.Body)
	fmt.Printf("Downloaded image size is %v\n", len(data))
}
