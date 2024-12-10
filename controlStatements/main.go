package main

import "fmt"

func main() {

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
