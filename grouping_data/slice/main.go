package main

import "fmt"

func main() {
	// Composite Literal
	x := []int {65, 66, 67, 68, 69, 70}

	// Slicing a slice
	fmt.Println(x[1:4])

	// Looping a slice using range
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
}