package main

import "fmt"

func main() {
	// Using Slicing with Slices
	x := []int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}