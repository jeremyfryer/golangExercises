package main

import "fmt"

func main(){
	fmt.Println(half(0))
	fmt.Println(half(1))
}

func half(num int) (int, bool) {
	return num/2, num%2 == 0
}