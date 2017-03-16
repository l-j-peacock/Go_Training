package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	// Casting (This will complain)
	fmt.Printf("%T\n", int(val))

	// Asserting (This will work)
	fmt.Printf("%T\n", val.(int))
}