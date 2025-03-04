package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple goroutine example
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Number %d\n", i)
	}
}

// Channel example
func calculateSquares(numbers chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		numbers <- i * i
	}
	close(numbers)
}

func main() {
	// Basic goroutine example
	go printNumbers()
	go printNumbers()

	// Using channels
	numbersChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	// Start goroutine with channel
	go calculateSquares(numbersChan, &wg)

	// Read from channel
	for num := range numbersChan {
		fmt.Printf("Received square: %d\n", num)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	time.Sleep(1 * time.Second)
}
