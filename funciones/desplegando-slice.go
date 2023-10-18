package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6}

	rp := foo("Hola", xi...) // convierte un slice de enteros con el operador de 3 puntos a un array de parametros que pueda entender la funcion
	/*
	 si se llama una funcion que espera n cnatidad de parametros y no se le pasa nada, por defecto su parametro sera nil
	*/
	fmt.Println(rp)
}

func foo(s string, x ...int) int { //... spread operator en JS. Indica que la cantidad de parametros en indefinida
	// go interpreta este operador ... como un array de valores de un tipo asignado. En este caso x ...int, x es un array de parametros de
	//tipo int

	/*
		este tipo de parametrp x ...int se debe usar al final en caso de que se reciban mas parametros de otros tipos

	*/
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(s)

	if x == nil {
		return 0
	}
	var suma int
	for i, valor := range x {
		fmt.Println(valor, i)

		suma += valor
	}

	return suma
}
