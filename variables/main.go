package main

import "fmt"

func main() {
	var name string = "My name is something"

	var isAdult bool = false

	var number int = 66

	floatingValue := 50

	println("The number is : ", number)
	println("The floatingValue is : ", floatingValue)

	if isAdult {
		fmt.Println("You are adult.")
	}

	fmt.Println(name)

	// Shorthand Syntax
	fullName := "Sushank ghimire"
	println(fullName)
}
