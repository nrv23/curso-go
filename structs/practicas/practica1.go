package main

import "fmt"

// declarar un tipo struct

type Persona struct { // funciona como una interfaz en JS
	name          string
	apelido       string
	edad          int
	saboresHelado []string
}

func main() {

	persona := Persona{
		name:    "Nataniel",
		apelido: "Venegas",
		edad:    32,
		saboresHelado: []string{
			"chocolate",
			"Vainilla",
			"crema",
			"galleta",
		},
	}

	fmt.Println(persona.name, persona.apelido, persona.edad)

	fmt.Println("Sabores favoritos")

	for key, sabor := range persona.saboresHelado {
		fmt.Println(key, sabor)
	}

}
