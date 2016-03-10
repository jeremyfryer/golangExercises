package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 1; fib(i) < 4000000; i++ {

		if fib(i)%2 == 0 {
			sum += fib(i)
		}
	}
	fmt.Println(sum)
}

func fib(x int) int {
	if x <= 1 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}

}
