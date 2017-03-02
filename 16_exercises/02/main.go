package main

import "fmt"

func main() {
	half := func(x int) {
		var isEven bool
		halfNum := x / 2
		if x%2 == 0 {
			isEven = true
		}
		fmt.Println("Result is:", halfNum, " , was the provided number even:", isEven)
	}

	half(3)
	half(5)
	half(7)
	half(9)
	half(11)
	half(2)
	half(4)
	half(6)
	half(8)
	half(10)
}