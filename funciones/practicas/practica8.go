package main

import "fmt"

func main() {

	fnSuma := suma(1, 2)
	fmt.Printf("Tipo de retorno %T \n", fnSuma)
	fmt.Printf("valor de retorno %v \n", fnSuma())
}

func suma(num1, num2 int) func() int {

	var suma = 0

	return func() int {
		suma = num1 + num2

		return suma
	}
}
