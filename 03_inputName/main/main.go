package main

import "fmt"

func main() {
	var n string
	fmt.Print("What is your name: ")
	fmt.Scan(&n)
	fmt.Printf("Hello %s", n)
}
