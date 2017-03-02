package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // Memory address

	changeMe(&age)

	fmt.Println(&age) // Same memory address
	fmt.Println(age) // 24
}

func changeMe(z *int) {
	fmt.Println(z) // Again, same memory address as above
	fmt.Println(*z) // 44
	*z = 24 // Reassignment of value at given memory address
	fmt.Println(z) // Again, same memory address as above
	fmt.Println(*z) // 24
}