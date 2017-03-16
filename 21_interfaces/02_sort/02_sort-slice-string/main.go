package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Before:", s)
	sort.Strings(s)
	fmt.Println("After:", s)
}

/* Another way to do it:

func main() {
	s := []string{"Zeno", "John", "Al" ,"Jenny"}
	fmt.Println(s)
	// sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
 */