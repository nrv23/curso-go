package main

import "fmt"

func main() {

	var x = []int{}
	var y = []int{}
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		//x = append(x, i, i, i, i) // retornar un slice del mismo tipo pero con el nuevo elemento agregado
		x = append(x, i)   // retornar un slice del mismo tipo pero con el nuevo elemento agregado
		y = append(y, i+2) // retornar un slice del mismo tipo pero con el nuevo elemento agregado
	}

	x = append(x, y...) // agregar un array dentro de otro
	/*
		y... el operador ... funciona como un spread operator de Javascript, toma el array y, cada elemnto lo va insertando en el array x
	*/
	fmt.Println(x)
	// eliminar elemntos de un slice
	// [0 1 2 3 4 5 6 7 8 9 2 3 4 5 6 7 8 9 10 11]
	x = append(x[:6], x[8:]...)
	/*
		sino se pone numero en la izquiera toma la primer posicion
		sino pone numero en la derecha entonces toma el length del array
	*/

	fmt.Println(x)
}
