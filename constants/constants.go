package main

import "fmt"

const age = 18

const (
	port = 3000
	host = "localhost:3000"
	frontend = "https://localhost:5173"
)

func main() {
	const name string = "Sushank ghimire"

	fmt.Println("Constants in Golang")

	fmt.Println("Your name is ", name, " and age is ", age)
}
