package main

import "fmt"

func main() {
	// Make a slice
	x := make([]int, 10, 12)

	// Composite Literal
	// x := []int {65, 66, 67, 68, 69, 70}

	// Appending to a slice
	// x = append(x, 71, 72, 73, 74, 75, 76)

	// Appending a slice to a slice
	// y := []int {71, 72, 73, 74, 75, 76}

	// x = append(x, y...)

	// Deleting from a slice
	// x = append(x[:5], x[7:]...)
	x[9] = 99
	x = append(x, 100, 101, 102)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))


	// Slicing a slice
	// fmt.Println(x[1:4])

	// Looping a slice using range
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
}