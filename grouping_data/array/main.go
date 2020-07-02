package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 65
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			fmt.Println(string(x[i]))
		}
	}
}