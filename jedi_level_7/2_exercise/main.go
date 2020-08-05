package main

import "fmt"

type person struct {
	first string
	surname string
}

func changeMe(p *person) {
	p.first = "Seun"
	// Or (*p).first = "Seun"
}

func main() {
	p1 := person {
		first: "Azeez",
		surname: "Banjoko",
	}

	changeMe(&p1)
	fmt.Println(p1)
}