package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex   sync.Mutex
)

func incrementCounter() {
	// Lock the mutex to ensure exclusive access to shared resource
	mutex.Lock()
	defer mutex.Unlock() // Ensure the mutex is unlocked when the function exits

	// Increment the counter
	counter++
	fmt.Println("Counter value:", counter)
}

func main() {
	// Create multiple goroutines to increment the counter concurrently
	for i := 0; i < 10; i++ {
		go incrementCounter()
	}

	// Sleep for a while to allow goroutines to complete
	// we can add wait grou as well here; but for now i added sleep
	time.Sleep(time.Second)
}
