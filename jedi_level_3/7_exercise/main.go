package main

import "fmt"

func main() {

	favSport := "Soccer"

	switch favSport {
	case "Soccer":
		fmt.Println("I love playing soccer")
	case "Basketball":
		fmt.Println("I love playing basketball")
	case "Cricket":
		fmt.Println("I love playing cricket")
	default:
		fmt.Println("I love other sports")
	}

}
