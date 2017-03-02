package main

import "fmt"

func main() {

	var myName string
	fmt.Println("Please enter your name:")
	fmt.Scan(&myName)
	fmt.Println("Hello", myName)
}
