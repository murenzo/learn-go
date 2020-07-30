package main

import "fmt"

func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func even(f func(num ...int) int, number ...int) int {
	var evenNumbers []int
	for _, v := range number {
		if v % 2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}

	return f(evenNumbers...)
}

func main() {
	numSlice := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(even(sum, numSlice...))
}