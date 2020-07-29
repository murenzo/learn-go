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

	m := map[string]Person {
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, v := range v.favorite_ice_cream_flavors {
			fmt.Println(i, v)
		}
		fmt.Println("-----------------")
	}
}
