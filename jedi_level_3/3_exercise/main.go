package main

import "fmt"

func main() {
	born_year := 1987
	for born_year <= 2020 {
		fmt.Println(born_year)
		born_year++
	}
}