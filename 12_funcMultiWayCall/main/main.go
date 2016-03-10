package main

import "fmt"

func main(){
	myFunc(1,2)
	myFunc(1,2,3)
	num_slice := []int{1, 5, 3, 7, 2}
	myFunc(num_slice...)
	myFunc()


}

func myFunc(num ...int) {
	fmt.Println(num)
}