package main

import "fmt"

func main() {
	isTrue := (true && false) || (false && true) || !(false && false)
	fmt.Println("The value is:", isTrue)
}