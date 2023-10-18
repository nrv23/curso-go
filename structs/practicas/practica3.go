package main

import "fmt"

type vehicle struct {
	puertas int
	color   string
}

type Camion struct {
	vehicle
	cuatroRuedas bool
}

type Sedan struct {
	vehicle
	lujoso bool
}

func main() {

	c1 := Camion{
		vehicle: vehicle{
			puertas: 4,
			color:   "negro",
		},
		cuatroRuedas: true,
	}

	s1 := Sedan{
		vehicle: vehicle{
			puertas: 2,
			color:   "blanco",
		},
		lujoso: true,
	}

	fmt.Println(c1)
	fmt.Println(s1)
	fmt.Println("")
	fmt.Println(c1.puertas)
	fmt.Println(s1.puertas)
}
