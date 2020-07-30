package main

import "fmt"

func main() {
	multiplier := func(ind int) func(n int) int {
		return func(num int) int {
			return ind * num
		}
	}

	multiplyBy2 := multiplier(2)

	for i := 1; i <= 12; i++ {
		fmt.Println(multiplyBy2(i))
	}
	
}