package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"Firstname":"Banjoko","Lastname":"Abdulazeez","Age":33},{"Firstname":"Raji","Lastname":"Sabith","Age":34}]`

	bs := []byte(s)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("The person index", i)
		fmt.Println(v.Firstname, v.Lastname, v.Age)
	}
}
