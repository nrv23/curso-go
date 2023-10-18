package main

import "fmt"

type Entero int

// declarar variables constantes
const s1 string = "Hola"
const s2 string = "Adios"
const v3 = 1
const v4 Entero = v3 // es posible si la constante v3 no se declara con el tipo sino se hace la conversion explicita

func main() {

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(v3)
	// imprimir los tipos
	fmt.Printf("%T \n", s1)
	fmt.Printf("%T \n", s2)
	fmt.Printf("%T \n", v3)
	fmt.Printf("%T \n", v4)
}
