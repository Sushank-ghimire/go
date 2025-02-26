package main

import "fmt"

func basics() {

	// For loop in golang
	for i := 0; i < 5; i++ {
		fmt.Println("Hello golang")
	}

	// Switch Statement in golang
	day := 7
	switch day {
	case 1:
		fmt.Println("It's sunday.")
		break
	case 2:
		fmt.Println("It's monday.")
		break

	default:
		fmt.Println("Invalid number.")
	}

	// If Else conditions in Golang
	age := 17
	if age < 18 {
		fmt.Println("You are a little children or underage")
	} else {
		fmt.Println("You are adult now.")
	}
}

func switchMulticases() {
	day := 7

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

func main() {
	// basics()
	switchMulticases()
}
