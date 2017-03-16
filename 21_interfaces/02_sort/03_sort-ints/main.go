package main

import (
"fmt"
"sort"
)

func main() {
	s := []int{1, 18, 24, 2, 7, 5, 1002, 45, 6}

	fmt.Println(s)
	sort.Sort(sort.IntSlice(s))
	//sort.Ints(s)
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

/* Could also do the following:

sort.Ints(s)

 */