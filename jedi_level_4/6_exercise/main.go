package main

import "fmt"

func main() {
	// Using make to create slices
	states := make([]string, 50, 50)
	states = []string {"Alabama", "Alaska", "Arizona", "Akansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii",
		"Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachussets", "Michigan",
		"Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", 
		"New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}