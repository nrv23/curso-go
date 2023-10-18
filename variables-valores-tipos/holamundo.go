package main

// si se importa un solo paquete seria asi import "fmt"
// importar varios paquetes

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, World!")
	nb, err := fmt.Println("Hola")

	if err == nil {
		fmt.Println("nb", nb)
	}

	_, _ = fmt.Println(nb, err) // le indica al compilador que no va usar los datos de retorno de la funcion
}
