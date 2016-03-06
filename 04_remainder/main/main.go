package main

import "fmt"

func main() {
	var small int
	var large int
	fmt.Print("What is the large number: ")
	fmt.Scan(&large)
	fmt.Print("What is the small number: ")
	fmt.Scan(&small)
	fmt.Printf("The remainder is %v", large%small)
}
