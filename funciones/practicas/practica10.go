package main

import "fmt"

func main() {
	// La función anónima 'closure' captura la variable 'x'.
	x := 10
	closure := func() {
		fmt.Println(x)
	}

	// Llamamos a la función anónima.
	closure() // Imprimirá "10".
}
