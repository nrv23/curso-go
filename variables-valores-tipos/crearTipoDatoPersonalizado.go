package main

import "fmt"

//crear un nuevo tipo de dato
type dinero int

var a int
var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("el tipo de dato de a es: %T \n", a)
	b = 1000
	fmt.Println(b)
	fmt.Printf("el tipo de dato de a es: %T \n", b)
}
