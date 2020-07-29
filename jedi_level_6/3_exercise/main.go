package main

import "fmt"

func main() {
	sum([]int {1,2,3,4,5,6,7,8,9}...)
}

func sum(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}

	defer message(total)
	return total
}

func message(n int) {
	fmt.Println("Total sum:", n)
}