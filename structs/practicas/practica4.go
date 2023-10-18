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
	// struct anonimo
	c1 := struct {
		vehicle struct {
			puertas int
			color   string
		}
		lujoso bool
	}{
		vehicle: struct {
			puertas int
			color   string
		}{
			puertas: 4,
			color:   "blanco",
		},
		lujoso: true,
	}

	fmt.Println(c1)
}
