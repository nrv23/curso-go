package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p Persona) presentarse() {

	fmt.Println("Hola mi nombre es " + p.Nombre + " " + p.Apellido)
	fmt.Printf("Mi edad es %v \n", p.Edad)
}

func main() {

	p1 := Persona{
		Nombre:   "Nataniel",
		Apellido: "Venegas",
		Edad:     31,
	}

	p1.presentarse()
}
