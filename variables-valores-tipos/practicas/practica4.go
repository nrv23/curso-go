package main

import "fmt"

type dinero int

var x dinero

func main() {

	fmt.Println(x)
	fmt.Printf("EL tipo de la variable x es %T\n", x)

	x = 42
	fmt.Println(x)

}
