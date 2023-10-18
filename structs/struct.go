package main

import "fmt"

// declarar un tipo struct

type Persona struct { // funciona como una interfaz en JS
	name string
	age  int
}

func main() {

	// declarar una variable de tipo Persona

	p1 := Persona{
		name: "Nataniel",
		age:  31,
	}

	p2 := Persona{
		name: "Nombre 2",
		age:  2,
	}

	fmt.Println(p1, p2)

	fmt.Println(p1.age)
	fmt.Println(p1.name)
	fmt.Println("")
	fmt.Println(p2.age)
	fmt.Println(p2.name)
}
