package main

import "fmt"

func basics() {
	var stringArray [2]string
	stringArray[0] = "first name"
	stringArray[1] = "second name"
	fmt.Println(stringArray)

	numArray := []int{10, 20, 30, 40, 50}

	for i, val := range numArray {
		fmt.Printf("The number index is %d and the value is %d\n", i, val)
	}
}

func GetUserInputNumbers() {
	var size int
	fmt.Println("Enter the size of the array : ")
	fmt.Scan(&size)
	fmt.Println("Size of the array : ", size)

	numArray := make([]int, size)

	for val := range numArray {
		fmt.Scan(&numArray[val])
	}

	fmt.Printf("The numbers are : ")
	for val := range numArray {
		fmt.Printf("%d", numArray[val])
	}
}

func main() {
	// basics()
	// GetUserInputNumbers()
	cars := []string{"Toyata", "BMW", "RRR"}
	for val := range cars {
		fmt.Println("Index : ", val, "Value : ", cars[val])
	}
}
