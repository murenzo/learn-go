package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		x := i % 4
		fmt.Printf("%d modulo 4: %d\n", i, x)
	}
}
