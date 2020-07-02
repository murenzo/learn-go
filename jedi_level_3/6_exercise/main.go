package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("This is a true story")
	case false:
		fmt.Println("This is a false story")
	default:
		fmt.Println("I do not know this kind of story")
	}
}
