package main

import "fmt"

func main() {
	n := []int {1,2,3,4,5,6,7,8,9}
	fmt.Println(bar(n))
}

func foo(num ...int)int {
	sum := 0
	for _, v := range num {
		sum += v
	}

	return sum
}

func bar(num []int)int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	
	return sum
}