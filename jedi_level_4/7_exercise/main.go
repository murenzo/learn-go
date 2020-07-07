package main

import "fmt"

func main()  {
	// Multi dimensional array
	a := []string {"James", "Bond", "Shaken not stirred"}

	b := []string {"Miss", "Moneypenny", "Hello, James"}

	c := [][]string {a, b}

	fmt.Println(c)

	for i1, v1 := range c {
		fmt.Printf("Inside array index %v\n", i1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}