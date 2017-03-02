package main

import "fmt"

func main() {

	var firstNumber int
	var secondNumber int
	var result int

	fmt.Println("Please enter a number below:")
	fmt.Scan(&firstNumber)
	fmt.Println("Please enter a second number below:")
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		result = firstNumber % secondNumber
	} else {
		result = secondNumber % firstNumber
	}

	fmt.Println("The result is:", result)
}
