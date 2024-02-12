package main

import "fmt"

func main() {
	fmt.Println("Channels in GoLang!")

	// Unbuffered Channel
	// this required publisher & Subscriber should running on same time
	UnbufferedChannelExErrored() // comment out this to run code
	fmt.Println("")
	UnbufferedChannelEx()

	// Buffered Channel
	// this can take input & store into buffer & consume when subscriber online
	fmt.Println("")
	BufferedChannelExErrored() // comment out this to run code
	fmt.Println("")
	BufferedChannelEx()
}

func UnbufferedChannelExErrored() {
	// push data into channel
	channelA := make(chan int)
	channelA <- 10

	// get data from channel
	channelData := <-channelA
	fmt.Println(channelData)
}

func UnbufferedChannelEx() {
	// push data into channel
	channelA := make(chan int)
	go func() { // started concurrent thred to achive result
		channelA <- 10
	}()

	// get data from channel
	channelData := <-channelA
	fmt.Println(channelData)
}

func BufferedChannelExErrored() {
	// push data int channel
	// in this case we have fixed buffer size 1
	// you can not push more than 1 int data into channel
	channelA := make(chan int, 1)
	channelA <- 10
	channelA <- 20

	// get data from channel
	channelData := <-channelA
	fmt.Println(channelData)
}

func BufferedChannelEx() {
	// push data int channel
	// in this case we have fixed buffer size 1
	// you can not push more than 1 int data into channel
	channelA := make(chan int, 2) // increased buffer size in channel
	channelA <- 10
	channelA <- 20

	// get data from channel
	channelData := <-channelA
	fmt.Println(channelData)
	channelData = <-channelA
	fmt.Println(channelData)
}
