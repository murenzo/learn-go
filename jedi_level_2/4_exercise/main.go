package main

// Bits Shifting 

import "fmt"

func main() {
	n := 4
	fmt.Printf("%d\t%b\t%#x\n", n, n, n)

	o := n << 1
	fmt.Printf("%d\t%b\t%#x\n", o, o, o)
}