package main

import "fmt"

func main() {

	// literal compuesto
	x := []int{1, 2, 3, 4, 5, 6, 7, 8} // un array que se declara con sus valores

	for _, elem := range x { // for range se comporta como un foreach. El primer valor es el indice, el segundo es el elemento iterado
		fmt.Println(elem)
		fmt.Printf("%T \n", elem)
	}
}
