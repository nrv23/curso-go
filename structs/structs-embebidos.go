package main

import "fmt"

// declarar un tipo struct

type Persona struct { // funciona como una interfaz en JS
	name string
	age  int
}

type agentesecreto struct {
	persona Persona // generar una estructura con valores de otra estructura
	lpm     bool
}

func main() {

	// declarar una variable de tipo Persona

	p1 := agentesecreto{

		persona: Persona{
			name: "Nataniel",
			age:  31,
		},
		lpm: true,
	}

	fmt.Println(p1.lpm)
	fmt.Println(p1.persona)
	fmt.Println(p1.persona.name)
	fmt.Println(p1.persona.age)
}
