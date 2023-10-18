package main

import "fmt"

func main() {
	resultado := foo(12, 3, 45, 56, 57, 568, 768, 789, 8, 98, 9)

	fmt.Println(resultado)
}

// func (r receptor) identificador(parametros) (retorno(s)) {codigo}
func foo(x ...int) int { //... spread operator en JS. Indica que la cantidad de parametros en indefinida
	// go interpreta este operador ... como un array de valores de un tipo asignado. En este caso x ...int, x es un array de parametros de
	//tipo int
	var suma int
	for i, valor := range x {
		fmt.Println(valor, i)

		suma += valor
	}

	return suma
}
