package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("My name is", s.fname, s.lname)
}

func main() {
	sa1 := secretAgent {
		person: person{
			fname: "Banjoko",
			lname: "Abdulazeez",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}
