package main

import "fmt"

func main() {
	st := foo()
	v := st()
	fmt.Println(v)
}

func foo() func() int {
	return func() int {
		return 33
	}
}