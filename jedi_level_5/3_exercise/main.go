package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck {
		vehicle: vehicle {
			doors: 2,
			color: "Brown",
		},
		fourWheel: true,
	}

	s1 := sedan {
		vehicle: vehicle {
			doors: 4,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(t1.vehicle.doors)
	fmt.Println(t1.vehicle.color)
	fmt.Println(t1.fourWheel)

	fmt.Println(s1.vehicle.doors)
	fmt.Println(s1.vehicle.color)
	fmt.Println(s1.luxury)
}