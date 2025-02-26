package main

import (
	"fmt"
)

type CarModels struct {
	price       int
	brand       string
	model       string
	stock       int
	releaseYear string
}

func main() {
	fmt.Print("-----Map in Golang-----\n")

	// Basic Map Declaration
	a := map[string]int{"hello": 4, "user": 7}
	fmt.Println(a)

	b := map[int]string{1: "one", 2: "Two", 3: "Three"}
	fmt.Println(b)

	// Using Make Function
	var carModel = make(map[string]string)

	carModel["price"] = "Rs. 3000000"
	carModel["brand"] = "Ford"
	carModel["model"] = "Mustang"
	carModel["year"] = "1964"

	// delete(carModel, "price_model")
	fmt.Println(carModel)

	for key, value := range carModel {
		fmt.Printf("\n%v : %v", key, value)
	}
}
