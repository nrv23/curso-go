package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a()) /// aumenta
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // apunta a otro espacio en memoria y cuenta de 0
	fmt.Println(b())
}

func incrementor() func() int {

	var x int // esto quieda fuera del scope, por lo que es invisble para el retorno

	return func() int { // esto se retorna y se mantiene el contador actualizado siempre y cuando apunte al mismo espacio en memoria
		// o a la misma variable que guarda la ejecucion
		/*
			Si la funcion incrementor se guarda en una variable y esa variable se llama n veces, las n veces el contador aumnta
			porque apunta al mismo espacio de memoria y cada vez que se llama, se retorna un ambiente pero con el contador aumentado
			Una vez que se guarda en otra variable, el espacio en memoria cambia y el contador se reinicia


		*/
		x++
		return x
	}
}
