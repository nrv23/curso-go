package main

import "fmt"

var a int
var b string = "programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "Dice Hola mundo" // string literal interpretado
	/*g := `EL doctor dice que comer vegetales
		es saludable
	` // este tipo de declaracion respeta el new line o los caracteres de escape */

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//fmt.Printf("El valor de la variable e es: %v  \n", e)
	//fmt.Printf("El valor de la variable b es: %v \n", b)
	fmt.Printf("El valor de la variable e es: %d \n", e)
	fmt.Printf("El valor de la variable b es: %s \n", b)
	// fmt.Println(f)
	// fmt.Println(g)

	//imprimir los tipos de datos en sinxtaxis GO

	fmt.Printf("El tipo de la variable a es: %T \n", a)
	fmt.Printf("El tipo de la variable a es: %T \n", b)

	//
	s1 := fmt.Sprint("El ", b, " ", f)
	fmt.Println(s1)
	fmt.Printf("El tipo de la variable s1 es: %T \n", s1)
}
