package main

import "fmt"

func main() {

	// literal compuesto
	x := []int{1, 2, 3, 4, 5, 6, 7, 8} // un array que se declara con sus valores

	/*for _, elem := range x { // for range se comporta como un foreach. El primer valor es el indice, el segundo es el elemento iterado
		fmt.Println(elem)
		fmt.Printf("%T \n", elem)
	} */

	fmt.Println(x[0:4]) // genera un tipo substring donde el valor de la izquierda es desde donde va obtener, y el numero de la derecha
	// hasta donde va obtener pero al numero de la derecha se le resta una posicion

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
