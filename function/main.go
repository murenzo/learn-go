package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello from function expression")
	}
	f()

	g := func(x int) {
		fmt.Println("My birth year is:", x)
	}
	g(1987)
	
}
