package main

import "fmt"

func main() {
	born_year := 1987
	for {
		if born_year > 2020 {
			break
		}
		fmt.Println(born_year)
		born_year++
	}
}
