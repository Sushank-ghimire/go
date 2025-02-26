package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var userGuess int
	var guessCouter int = 0
	targetNum := rand.Intn(100) + 1 // Random number between (1-100)

	fmt.Println("-------Number Guessing Game----------")

	for {
		fmt.Printf("Enter your guess number (%v) : ", targetNum)
		fmt.Scan(&userGuess)
		guessCouter++
		if userGuess < targetNum {
			fmt.Println("Your guess is less than that of original number")
		} else if userGuess > targetNum {
			fmt.Println("Your guess is less than that of original number")
		} else {
			fmt.Printf("Your won. Correct guess in %v times", guessCouter)
			break
		}
	}
}
