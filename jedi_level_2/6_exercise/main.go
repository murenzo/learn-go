package main

// iota

import "fmt"

const (
	a = iota + 2016
	b
	c
	d

)

func main() {
	fmt.Println(a, b, c, d)
}