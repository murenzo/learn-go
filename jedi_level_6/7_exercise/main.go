package main

import "fmt"

var sum func(n int, o int) int

func main() {
	sum = func (a int, b int) int {
		return a + b
	}

	fmt.Println(sum(2, 3))
}
