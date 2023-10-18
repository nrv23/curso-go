package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a) // con & se puede imprimir el punterpo en memoria

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", &a) // con & se puede imprimir el punterpo en memoria

	var b *int = &a //asginar un valor de mempria a otra variable

	fmt.Printf("%v \n", b)

	// acceder al valor que esta almacrnado en el putero
	// con el * se imprime el valor del puntero en memoria
	fmt.Println(*b)

	// setear el valor al puntero

	*b = 100

	fmt.Println(*b)
}
