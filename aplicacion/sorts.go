package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println("")
	sort.Ints(xi)    // ordenar array de numerps
	sort.Strings(xs) // ordenar array de strings
	fmt.Println("")
	fmt.Println(xi)
	fmt.Println(xs)
	/*

		las funciones de ordenamiento no tienen retorno, manejan el valor internamento dentro del puntero
	*/
}
