package main

import "fmt"

var x int = 2
var z int = 10

func main() {
	numero(x)
	numero2()
	x := 2 + 7 // el scope de la funcion se sobrepone sobre el scope a nivel de paquete
	y := "james bond"

	fmt.Println("valor interno ", x)
	fmt.Println(y)
	x = 5
	fmt.Println("valor interno ", x)

}

/*
	primer parentesis son los parametros de la funcion
	segundo parentesis, es opcional si la funcion retorna algo, debe escribir los valores y tipos de retorno en el mismo orden
	en que se va retornando
*/
func numero(x int) /*(n int, err error) */ {
	fmt.Println("valor externo ", x)
}

func numero2() {
	fmt.Println(z)
}
