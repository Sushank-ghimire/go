package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditions In Golang")

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}
}
