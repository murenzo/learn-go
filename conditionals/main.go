package main

import "fmt"

func main() {
	x := 422
	if x == 40 {
		fmt.Println("our number is 40")
	} else if x == 41 {
		fmt.Println("our number is 41")
	} else if x == 42 {
		fmt.Println("our number is 42")
	} else {
		fmt.Println("our number is not 40, 41, 42")
	}
}
