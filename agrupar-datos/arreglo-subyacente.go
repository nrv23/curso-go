package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	//y := []int{}
	z := []int{}

	z = append(x[:2], x[3:]...)
	/*y = append(x, 6, 7, 8, 9) /* // devuelve el arreglo original x por no superar su capacidad
	El arreglo x tiene una capacidad de 5, pero como al agregar mas elementos se supera la capacidad de x, se gener aun nuevo slice
	que es retornado por append con los nuevos valores y memoria dinamica asignada por go automaticamente
	Si la capacidad del arreglo original nos e supera entonces se retorna el mismo arreglo con los elementos actualzados
	*/

	fmt.Println(x)
	//fmt.Println(y) // y guarda el nuevo arreglo
	fmt.Println(x) // y guarda el nuevo arreglo
	fmt.Println(z) // y guarda el nuevo arreglo
}
