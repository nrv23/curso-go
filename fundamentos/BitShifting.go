package main

import "fmt"

const (
	_  = iota             // desecha el primer valor del iota que es el 0
	kb = 1 << (iota * 10) // iota vale 1 y se multiplica por 10 para indicar que haga el bigshifting a la izquierda 10 veces
	gb = 1 << (iota * 10) // iota vale 2
	tb = 1 << (iota * 10) // iota vale 3
)

func main() {

	a := 4
	fmt.Println("Valor original")
	fmt.Printf("%d \t\t %b", a, a)
	fmt.Println("")
	// %b indica formato binario

	b := a << 2 //hacer bigshifting a la izquierda
	fmt.Println("bigshift izquierda")
	fmt.Println("")
	fmt.Printf("%d \t\t %b", b, b)
	fmt.Println("")
	b = a >> 2 //hacer bigshifting a la derecha

	fmt.Println("bigshift izquierda")
	fmt.Println("")
	fmt.Printf("%d \t\t %b", b, b)

	// ejemplos

	fmt.Println("")
	fmt.Println("")

	fmt.Printf("%d \t\t %b", kb, kb)
	fmt.Println("")
	fmt.Printf("%d \t %b", gb, gb)
	fmt.Println("")
	fmt.Printf("%d \t %b", tb, tb)
	fmt.Println("")
}
