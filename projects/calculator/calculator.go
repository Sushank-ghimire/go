package main

import (
	"fmt"
)

func Calculator(a float64, b float64, operator string) {
	switch operator {
	case "-":
		fmt.Println("Your result of num1 - num2 is : ", a-b)
		break
	case "*":
		fmt.Println("Your result of num1 * num2 is : ", a*b)
		break

	case "+":
		fmt.Println("Your result of num1 + num2 is : ", a+b)
		break
	case "/":
		fmt.Printf("Your result of num1 / num2 is %.2f: ", a/b)
		break
	default:
		fmt.Printf("\nInvalid Operator or data found")
	}
}

func main() {
	var num1, num2 float64
	var operator string
	fmt.Print("Enter two numbers for the calculation : ")
	fmt.Scan(&num1, &num2)

	fmt.Print("Enter the operator (*, +, -, /): ")
	fmt.Scan(&operator)

	Calculator(num1, num2, operator)
}
