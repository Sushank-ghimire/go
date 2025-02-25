package main

import (
	"fmt"
)

func main() {
	// Boolean
	var isLoggedIn bool = false
	fmt.Println("Is user logged in : ", isLoggedIn)

	// Integers used for Whole Numbers (int, int8, int16, int32, int64)
	var num int = 5
	var num1 int8 = 127
	var num2 int64 = 48448484

	fmt.Println("Num : ", num)
	fmt.Println("Medium num : ", num1)
	fmt.Println("Large num : ", num2)

	// Unsigned Integer Stores only positive integers (uint, uint8, uint16, uint32, uint64)
	var num3 uint = 44
	var num4 uint64 = 444444
	fmt.Println("Unsigned integer : ", num3)
	fmt.Println("Unsigned 64 bit integer : ", num4)

	// Floating Points (float32, float64)
	var price float32 = 44.34
	fmt.Println("Floating Price : ", price)

	// Complex Numbers
	var c1 complex64 = complex(4, 5)
	c1 = complex(5, 7)
	fmt.Println("Complex Number : ", c1)

	// Array type
	var arr [3]int = [3]int{1, 2, 3}
	for val := range arr {
		fmt.Println("Values of Array : ", val)
	}

	// Slice for dynamic resizable
	slice := []int{1, 2, 3, 4, 5, 6}
	for val := range slice {
		fmt.Println("Slice Value Array : ", val)
	}
}
