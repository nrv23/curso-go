package main

import "fmt"

func main() {

	func() { // funciona anomia que se autollama
		fmt.Println("Se ejecutó la funcion anonima automaticamente")
	}() // al poner los parentesis la funcion se llama

	func(param int) { // funciona anomia que se autollama
		fmt.Printf("Argumento de tipo %T \n", param)
		fmt.Printf("Valor del argumento %v \n", param)
	}(24) // al poner los parentesis la funcion se llama
}
