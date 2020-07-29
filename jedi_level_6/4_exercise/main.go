package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "& I am", p.age, "years old")
}

func main() {
	p1 := person {
		first: "Banjoko",
		last: "Abdulazeez",
		age: 33,
	}

	p1.speak()
}