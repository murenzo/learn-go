package main

import (
	"fmt" 
	"encoding/json"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	p1 := person{
		Firstname: "Banjoko",
		Lastname:  "Abdulazeez",
		Age:       33,
	}

	p2 := person{
		Firstname: "Raji",
		Lastname:  "Sabith",
		Age:       34,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
