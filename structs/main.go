package main

import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

func main() {
	// Anonymous Struct
	p1 := struct {
		first string
		last string
		age int
	} {
		first: "Banjoko",
		last: "Abdulazeez",
		age: 33,
	}
	// sa1 := secretAgent {
	// 	person: person{
	// 		first: "Banjoko",
	// 		last:  "Abdulazeez",
	// 		age:   33,
	// 	},
	// 	ltk: true,
	// }

	// p2 := person{
	// 	first: "Raji-Banjoko",
	// 	last:  "Sabith",
	// 	age:   34,
	// }

	fmt.Println(p1.first, p1.last, p1.age)
}
