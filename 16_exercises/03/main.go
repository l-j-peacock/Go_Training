package main

import "fmt"

func main() {
	fmt.Println(greatestNumberFinder(1, 13, 65, 21, 34, 99, 20))
}

func greatestNumberFinder(x ...int) int {
	var largestNum int
	for _, v := range x {
		if v > largestNum {
			largestNum = v
		}
	}
	return largestNum
}