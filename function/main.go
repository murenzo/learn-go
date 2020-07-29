package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	foo(xi...)
}

// Variadic parameters
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
		fmt.Println("The number", v, "will sum up to the total of", sum)
	}
	fmt.Println("The total sum is", sum)
}

