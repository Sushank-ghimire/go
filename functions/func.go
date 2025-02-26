package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func parameterAcceptor(name string, age int) {
	fmt.Printf("Your name is : %s and your age is : %d", name, age)
}

func GreetUser(name string) {
	fmt.Printf("Hello, %s", name)
}

func FunctionReturn(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("------Functions in golang------")
	myMessage()
	parameterAcceptor("Sushank Ghimire", 18)
}
