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

// Todd's Solution is below:

/*
func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
 */