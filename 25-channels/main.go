package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Channels in GoLang!")

	// Unbuffered Channel
	// this required publisher & Subscriber should running on same time
	// UnbufferedChannelExErrored() // comment out this to run code
	fmt.Println("")
	UnbufferedChannelEx()

	// Buffered Channel
	// this can take input & store into buffer & consume when subscriber online
	fmt.Println("")
	// BufferedChannelExErrored() // comment out this to run code
	fmt.Println("")
	BufferedChannelEx()

	// One more some realistic unbuffered channel example
	RealisticUnbufferedChannelEx()
	fmt.Println("")
	time.Sleep(2 * time.Second)
	RealisticBufferedChannelEx()
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

func getMyLuckyNumber() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func RealisticUnbufferedChannelEx() {
	relChan := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				num := getMyLuckyNumber()
				relChan <- num
			}()
		}
		wg.Wait()
		close(relChan)
	}()

	for data := range relChan {
		fmt.Println(data)
	}
}

func RealisticBufferedChannelEx() {
	bufferLength := 1000
	relChan := make(chan int, bufferLength)
	wg := sync.WaitGroup{}

	for i := 0; i < bufferLength; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := getMyLuckyNumber()
			relChan <- num
		}()
	}
	wg.Wait()
	close(relChan)

	for data := range relChan {
		fmt.Println(data)
	}
}
