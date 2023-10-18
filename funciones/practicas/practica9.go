package main

import "fmt"

func main() {
	fmt.Println(sumaDos(suma, 1, 2, 3, 34, 3, 434, 45, 645, 654, 54, 645, 4, 5, 5656, 3, 2, 3, 2))
	fmt.Println(sumatres(suma, 1, 2, 3, 34, 3, 434, 45, 645, 654, 54, 645, 4, 5, 5656, 3, 2, 3, 2))
}

func suma(a ...int) int { // esta funcion es el callback de retorno

	var suma int = 0

	for _, elemento := range a {

		suma += elemento
	}

	return suma
}

func sumaDos(fn func(i ...int) int, b ...int) int {

	var suma []int

	for _, elemento := range b {

		if elemento == 2 {
			suma = append(suma, elemento)
		}
	}

	return fn(suma...)
}

func sumatres(fn func(i ...int) int, b ...int) int {

	var suma []int

	for _, elemento := range b {

		if elemento == 3 {
			suma = append(suma, elemento)
		}
	}

	return fn(suma...)
}
