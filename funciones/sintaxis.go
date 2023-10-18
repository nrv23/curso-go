package main

import "fmt"

func main() {
	foo()
	bar("Sdasd")
	rp := woo("Nataniel")
	saludo, estado := saludar("Nataniel", "Venegas")
	fmt.Println(rp)
	fmt.Println(saludo, estado)
}

func foo() {
	fmt.Println("Hola desde foo")
}

func bar(s string) {
	fmt.Println(s)
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo ,", st)
}

// como nombre y apellido son del mismo tipo se pone el tipo al final
// tambien se puede delcarar nombre string, apellido string
func saludar(nombre, apellido string) (string, bool) { // las funciones en go pueden tener varios retornos al mismo tiempo
	p := fmt.Sprint(nombre, " ", apellido, `dice "Hola"`)
	q := true

	return p, q
}
