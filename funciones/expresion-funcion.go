package main

import "fmt"

// declarar una funcion dentro de una variable en el scope global
var funcAnomina = func (i string ) => i

func main() {

	func1 := func(n1 int, n2 int) int { // funcion anomima que se guarda en una variable

		return n1 + n2
	}

	fmt.Println(func1(1, 2))

}

