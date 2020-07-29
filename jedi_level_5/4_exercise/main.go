package main

import "fmt"

func main() {
	p := struct {
		attributes map[string]string
		favColors []string
	} {
		attributes: map[string]string {"height": "6ft", "color": "Brown", "bestFood": "Plantain & Egg",},
		favColors: []string {"Black", "Brown", "Ash"},
	}

	fmt.Println(p)
}