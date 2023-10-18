package main

import "fmt"

// declarar una funcion dentro de una variable en el scope global
var funcAnomina = func(s string, s2 string) string {

	return s + " " + s2
}

func main() {

	foo()
	fmt.Println("")

	func1 := func(n1 int, n2 int) int { // funcion anomima que se guarda en una variable

		return n1 + n2
	}

	fmt.Println(func1(1, 2))

	func() { // funciona anomia que se autollama
		fmt.Println("Se ejecutó la funcion anonima automaticamente")
	}() // al poner los parentesis la funcion se llama

	func(param int) { // funciona anomia que se autollama
		fmt.Println("imprimiendo argumento ", param)
	}(24) // al poner los parentesis la funcion se llama

	fmt.Println(funcAnomina("Hola", "Mundo"))
}

func foo() {
	fmt.Println("Se ejecutó la funcion")
}
