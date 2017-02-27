package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
