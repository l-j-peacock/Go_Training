package main

import "fmt"

func half(x int) (string, string) {
	var isEven bool
	halfNum := x / 2
	if x%2 == 0 {
		isEven = true
	}
	return fmt.Sprint("Result is:", halfNum), fmt.Sprint("Was it even:", isEven)
}

func main() {
	fmt.Println(half(15))
	fmt.Println(half(10))
}