package main

import "fmt"

func main() {
	foo(1,2,3,4,5,6,7,8,9)
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

// func foo() {
// 	fmt.Println("Hello from foo baba")
// }

// func bar(s string) {
// 	fmt.Println("Hello", s)
// }

// func baz(s string) string {
// 	return fmt.Sprint("Hello from woo, ", s)
// }

// func mouse(fn string, ln string) (string, bool) {
// 	s := fmt.Sprint(fn, " " ,ln, ` says "hello"`)
// 	b := false
// 	return s, b
// }
