package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 243, 454, 534, 6, 456}

	fmt.Println(suma(ii...))

	fmt.Println("resultado de la suma de pares ", SumaPares(suma, ii...))
	fmt.Println("resultado de la suma de impares ", SumaImPares(suma, ii...))
}

func suma(x ...int) int {

	total := 0

	for _, num := range x {

		total += num
	}

	return total
}

// declaracion de un callback

func SumaPares(f func(x ...int) int, vi ...int) int { // funcion que retorna un call back
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v) // un array de numeros pares
		}
	}

	return f(y...) // obtiene la suma total de todos los pares
}

func SumaImPares(f func(x ...int) int, vi ...int) int { // funcion que retorna un call back
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v) // un array de numeros pares
		}
	}

	return f(y...) // obtiene la suma total de todos los pares
}
