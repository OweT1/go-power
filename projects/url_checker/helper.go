package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Helper functions
func checkUrl(url string) {
	_, err := http.Get(url)
	if err == nil {
		fmt.Printf("%s is up! :)\n", url)
	} else {
		fmt.Printf("%s is down! :(\n", url)
	}
}

func checkUrlChannel(url string, c chan string) {
	_, err := http.Get(url)
	if err == nil {
		c <- fmt.Sprintf("%s is up! :)", url)
	} else {
		c <- fmt.Sprintf("%s is down! :(", url)
	}
}

// Checker functions
func linearCheck(websites []string) {
	// Linear way
	fmt.Println("Linear Method")
	start := time.Now()
	for _, url := range websites {
		fmt.Printf("Checking %s...\n", url)
		checkUrl(url)
	}
	time_taken := time.Since(start)
	fmt.Printf("Linear Method - Total Time taken: %s\n", time_taken)
}

func concurrentCheck(websites []string) {
	// Concurrent way
	fmt.Println("Concurrent Method | Sync WaitGroup")
	var wg sync.WaitGroup

	start := time.Now()
	for _, url := range websites {
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

	time_taken := time.Since(start)
	fmt.Printf("Concurrent Method | Sync WaitGroup - Total Time taken: %s\n", time_taken)
}

func concurrentCheckChannel(websites []string) {
	// Concurrent way
	fmt.Println("Concurrent Method | Channel")
	c := make(chan string)

	start := time.Now()
	for _, url := range websites {
		fmt.Printf("Checking %s...\n", url)
		go checkUrlChannel(url, c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}

	time_taken := time.Since(start)
	fmt.Printf("Concurrent Method | Channel - Total Time taken: %s\n", time_taken)
}