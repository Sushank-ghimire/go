package main

import (
	"fmt"
)

func SliceDeclaration() {

	// String Slice
	strSlice := []string{}
	fmt.Println("String Slice : ", strSlice)

	// Boolean Slice
	boolSlice := []bool{true, false, true, false}
	fmt.Println("String Slice : ", boolSlice)

	// Integer slice
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("String Slice : ", intSlice)

	// Floating Value Slice
	floatSlice := []float32{33.33, 334.532, 5.5, 6.6}
	fmt.Println("String Slice : ", floatSlice)

	// Slice Methods
	fmt.Println(len(strSlice))
	fmt.Println(cap(strSlice))

	// My Slice From previously declared Array
	mySlice := floatSlice[1:3]
	fmt.Println("My Slice : ", mySlice)

	// Make Function
	mySlice2 := make([]int, 4)
	fmt.Println("Myslice 2 Value : ", mySlice2)
	fmt.Println("Myslice 2 Capacity : ", cap(mySlice2))
	fmt.Println("Myslice 2 Length : ", len(mySlice2))
}

func main() {
	fmt.Printf("Slices in golang \n")
	// SliceDeclaration()

	// Slice Modifications
	modifyingSlice := []string{}
	modifyingSlice = append(modifyingSlice, "Hello_world", "added_value")
	fmt.Println("Modified Slice : ", modifyingSlice, "Length : ", len(modifyingSlice))

	// Slices appending
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}

	mySlice3 := append(myslice1, myslice2...)
	fmt.Println("Appended new slice : ", mySlice3)

	// Copy method
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	neededNumbers := numbers[:len(numbers)-10]

	numbersCopy := make([]int, len(neededNumbers))
	
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
