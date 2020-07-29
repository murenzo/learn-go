package main

import "fmt"

func main() {
	fac := loopFact(4)
	fmt.Println(fac)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	factorialSum := 1
	for ; n > 0; n-- {
		factorialSum *= n
	}
	return factorialSum
}
