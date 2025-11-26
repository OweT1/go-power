package main

import (
	"fmt"
)

func logResults(answered, correct, num_problems int) {
	fmt.Printf("You answered %d out of %d questions.\n", answered, num_problems)
	fmt.Printf("You scored %d out of %d questions.\n", correct, num_problems)
}