package main

import (
	"fmt"
	"github.com/l-j-peacock/Go_Training/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}