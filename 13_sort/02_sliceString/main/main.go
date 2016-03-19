package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)

	fmt.Println(s)

	sr := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(sr)))
	fmt.Println(sr)

	mr := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(MyReverse(sort.StringSlice(mr)))
	fmt.Println(mr)

}

type myReverse struct {
	sort.Interface
}

func (r myReverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func MyReverse(data sort.Interface) sort.Interface {
	return &myReverse{data}
}
