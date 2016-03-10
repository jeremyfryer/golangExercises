package main

import "fmt"

func main() {
	num_slice := []int{1, 5, 3, 7, 2}
	fmt.Println(greatest(num_slice...))

}

func greatest(num ...int) int {
	total := 0
	for _, v := range num {
		if v > total {
			total = v

		}
	}
	return total
}
