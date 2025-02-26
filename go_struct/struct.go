package main

import (
	"fmt"
)

type Person struct {
	name      string
	id        int
	age       int
	working   string
	isMarried bool
}

func main() {
	user := Person{}

	fmt.Print("Enter Name: ")
	fmt.Scanln(&user.name) // Using Scanln to read string input properly

	fmt.Print("Enter Age: ")
	fmt.Scan(&user.age) // Using Scan to read integer input

	fmt.Print("Enter your ID : ")
	fmt.Scan(&user.id)

	fmt.Print("Enter marital status (true/false) : ")
	fmt.Scan(&user.isMarried)

	fmt.Println("\nUser Details:")
	fmt.Println("Name :", user.name)
	fmt.Println("Age :", user.age)
	fmt.Println("User ID :", user.id)
	fmt.Println("Is married :", user.isMarried)
}
