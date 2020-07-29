package main

import "fmt"

func main() {
	age := foo()
	myAge, myName := bar()

	fmt.Println(age)
	fmt.Println(myAge)
	fmt.Println(myName)
}

func foo()int {
	return 23
}

func bar() (int, string) {
	age := 34
	name := "Banjoko Abdulazeez"

	return age, name
}