package main

import "fmt"

type murenzo int

var x murenzo

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}