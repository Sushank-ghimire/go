package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("--------File Reader----------")
	fmt.Printf("Enter the filename to read : ")
	var filename string
	fmt.Scan(&filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:\n", string(data))
}
