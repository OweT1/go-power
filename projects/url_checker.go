package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Constants
var WEBSITES = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://fake-url.com",
}

// Helper functions
func checkUrl(url string) {
	_, err := http.Get(url)
	if err == nil {
		fmt.Printf("%s is up! :)\n", url)
	} else {
		fmt.Printf("%s is down! :(\n", url)
	}
}

// Main
func main() {
	// Linear way
	fmt.Println("Linear Method")
	start := time.Now()
	for _, url := range WEBSITES {
		fmt.Printf("Checking %s...\n", url)
		checkUrl(url)
	}
	time_taken := time.Since(start)
	fmt.Printf("Linear Method - Total Time taken: %s\n", time_taken)

	// Concurrent way
	fmt.Println("Concurrent Method")
	var wg sync.WaitGroup

	start = time.Now()
	for _, url := range WEBSITES {
		fmt.Printf("Checking %s...\n", url)
		wg.Add(1)

		// Defining of Goroutine
		go func(diff_url string) {
			defer wg.Done() // This statement will only execute after all the code in the rest of the function finishes executing
			checkUrl(diff_url)
		}(url)
	}

	// Block till counter reaches 0
	wg.Wait()

	time_taken = time.Since(start)
	fmt.Printf("Concurrent Method - Total Time taken: %s\n", time_taken)
}