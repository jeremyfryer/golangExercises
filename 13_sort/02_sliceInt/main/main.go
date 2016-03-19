package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)

	nr := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(nr)))
	fmt.Println(nr)
}
