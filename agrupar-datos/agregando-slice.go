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
}
