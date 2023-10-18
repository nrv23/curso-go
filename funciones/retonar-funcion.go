package main

import "fmt"

// declarar una funcion dentro de una variable en el scope global

func main() {

	fmt.Printf("%T\n", foo(1, 2)) // el tipo de retorno
	fmt.Println(foo(1, 2)())      // valor de retorno
	/*
		una funcion que retorna una funcion debe llamar la funcion original mas los parentesis de la funcion que retorna
		ejemplo: foo(1, 2)()
	*/
}

// funcion que retorna una funcion

func foo(var1 int, var2 int) func() int {
	return func() int {
		return var1 + var2
	}
}
