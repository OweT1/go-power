package main

import (
	"fmt"
	"os"
	"quiz_console/utils"
	"strings"
	"time"
)

// Constants
var FILENAME string = "data/problems.csv"
var TIMELIMIT int = 30

func logResults(answered, correct, num_problems int) {
	fmt.Printf("You answered %d out of %d questions.\n", answered, num_problems)
	fmt.Printf("You scored %d out of %d questions.\n", correct, num_problems)
}

func main() {
	file, err := utils.OpenFile(FILENAME)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", FILENAME)
		os.Exit(1)
	}

	// File opened successfully
	fmt.Printf("File `%s` opened successfully!\n", FILENAME)
	defer file.Close() // to close the file once we are done with the code

	// Parse file as CSV
	problems := utils.ParseCsv(file)

	// Initialise Quiz
	timer := time.NewTimer(time.Duration(TIMELIMIT) * time.Second)
	correct, answered := 0, 0

	// Create channel for user's answer
	answerCh := make(chan string)
	

	// Initialise Game Loop as a Label, to be broken later on
	GameLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Question)

		// Goroutine to listen for user's input in the background
		go func() {
			var answer string
			// Scanf reads user input into the variable 'answer'
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// Selects the channel that comes first
		select {
		case <- timer.C:
			// If the Timer channel sends a signal (time is up)
			fmt.Printf("\nTime out!\n")
			break GameLoop

		case answer := <- answerCh:
			if strings.ToLower(answer) == strings.ToLower(p.Answer) {
				correct++
			}
			answered++
		}
	}
	logResults(answered, correct, len(problems))
}

