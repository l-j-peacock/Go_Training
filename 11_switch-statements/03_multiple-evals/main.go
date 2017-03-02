package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Wassup Marcus or Medhi")
	case "Juilian", "Sushant":
		fmt.Println("Wassup Juilian or Sushant")
	}
}
