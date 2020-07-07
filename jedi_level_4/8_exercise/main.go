package main

import "fmt"

func main() {
	// Using map to create key value pair data structure
	persons := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dir":          []string{"Being Evil", "Ice cream", "Sunsets"},
	}

	// Adding value to the map
	persons["azeez"] = []string {"Family", "Worshiping Allah", "Coding"}

	// Deleting key value from map
	if _, ok := persons["no_dir"]; ok {
		delete(persons, "no_dir")
	}

	for k, fav := range persons {
		fmt.Println("This is the value for key ", k)
		for i, v := range fav {
			fmt.Println("\t", i, v)
		}
	}
	// fmt.Println(persons)
}
