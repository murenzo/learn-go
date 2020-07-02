package main

import "fmt"

func main() {
	// for x := 0; x < 100; {
	// 	x++
	// 	if x%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(x)
	// }
	// fmt.Println("Done")

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
