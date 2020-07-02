package main

import "fmt"

func main() {
	// Composite Literal
	x := []int {65, 66, 67, 68, 69, 70}

	// Appending to a slice
	// x = append(x, 71, 72, 73, 74, 75, 76)

	// Appending a slice to a slice
	y := []int {71, 72, 73, 74, 75, 76}

	x = append(x, y...)

	fmt.Println(x)


	// Slicing a slice
	// fmt.Println(x[1:4])

	// Looping a slice using range
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
}