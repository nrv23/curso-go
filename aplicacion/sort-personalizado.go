package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type ByAge []Persona
type ByName []Persona

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

// ordenar por nombre

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {
	p1 := Persona{"Eduar", 32}
	p2 := Persona{"María", 25}
	p3 := Persona{"Carolina", 56}
	p4 := Persona{"Adriana", 60}

	personas := []Persona{p1, p2, p3, p4}

	fmt.Println(personas)

	sort.Sort(ByAge(personas)) // esta funcion sort recibe un parametro de tipo Interface y cualquier valor que se le pase
	// como parametro debe satisface lo que el tipo de valor Interface tenga
	// una vez que se satusface la interfaz, la funcion ejecuta el metodo implicitamente
	fmt.Println(personas)
	sort.Sort(ByName(personas))
	fmt.Println(personas)
	// ordenamiento de datos de forma personalizada

	// ordena el valor a nivel de valor de puntero, por eso no retorna nada
	/*
		para que se pueda ordenar un array de datos de forma personalizada, el tipo de dato debe satisfacer la interfaz
		Interface que tiene metodos swap,len y less. Entonces si un valor de x interfaz puede implementar el metodo swap,
		quiere decir que ese valor implementa la interfaz Interface
	*/

}
