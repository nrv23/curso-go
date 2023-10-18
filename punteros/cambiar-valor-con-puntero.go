package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

/*
al pasar un puntero como parametro, tanto el argumento de la funcion como el parametro de la funcion
acceden al mismo valor en memoria y por lo tanto su valor en memoria es el mismo y al cambiar uno de los dos
el cambio se replica en ambos
*/
func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
