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

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("My name is", p.fname, p.lname, "says person")
}

func (s secretAgent) speak() {
	fmt.Println("My name is", s.fname, s.lname, "says secret agent")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Human was passed into bar", h.(person).fname)
	case secretAgent:
		fmt.Println("Secret Agent was passed into bar", h.(secretAgent).fname)
	default:
	fmt.Println("I was passed into bar", h)
	}
	
}

func main() {
	sa1 := secretAgent {
		person: person{
			fname: "Banjoko",
			lname: "Abdulazeez",
		},
		ltk: true,
	}

	p1 := person {
		fname: "Raji-Banjoko",
		lname: "Sabith",
	}

	bar(sa1)
	bar(p1)
}
