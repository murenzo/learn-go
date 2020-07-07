package main

import "fmt"

func main() {
	m := map[string]int{
		"Azeez":  33,
		"Bose":   34,
		"Nailah": 5,
		"Faiz":   3,
	}

	m["Lateef"] = 31
	m["Kazeem"] = 28

	for k, v := range m {
		fmt.Println(k, v)
	}

	if _, ok := m["Lateef"]; ok {
		delete(m, "Lateef")
	}

	fmt.Println(m)
	// if _, ok := m["Bose"]; ok {
	// 	fmt.Println("This is the print line inside the if statement")
	// }
	// fmt.Println(m["Azeez"])
}
