package main

import "fmt"

func main() {
	increment := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}

	zer := increment()
	fmt.Println(zer())
	fmt.Println(zer())
	fmt.Println(zer())
	fmt.Println(zer())
	fmt.Println(zer())
}