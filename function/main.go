package main

import "fmt"

func main() {
	n := []int {1,2,3,4,5,6,7,8,9,10}
	sumOddTotal := odd(sum, n...)
	fmt.Println(sumOddTotal)
}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, xx ...int) int {
	var yi []int
	for _, v := range xx {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func odd(f func(n ...int) int, v ...int) int {
	var oddNumbers []int
	for _, v := range v {
		if v % 2 != 0 {
			oddNumbers = append(oddNumbers, v)
		}
	}

	return f(oddNumbers...)
}