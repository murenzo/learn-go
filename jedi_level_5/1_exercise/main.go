package main

import "fmt"

type Person struct {
	first_name                 string
	last_name                  string
	favorite_ice_cream_flavors []string
}

func main() {
	p1 := Person{
		first_name:                 "Banjoko",
		last_name:                  "Abdulazeez",
		favorite_ice_cream_flavors: []string{"Chocolate", "Strawberry", "Raspberry"},
	}

	p2 := Person{
		first_name:                 "Raji-Banjoko",
		last_name:                  "Sabith",
		favorite_ice_cream_flavors: []string{"Strawberry", "Chocolate"},
	}
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)

	for i, v := range p1.favorite_ice_cream_flavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)

	for i, v := range p2.favorite_ice_cream_flavors {
		fmt.Println(i, v)
	}
}
