package main

import "fmt"

func main() {

	m1 := map[string]int{ //el nombre de la llave va en string y el valor en int
		"batman": 32,
		"robin":  40,
	}

	fmt.Println(m1)

	// leer propiedad del mapa

	//val, ok := m1["indice"]

	if val, ok := m1["indice"]; ok { // dentro del if se asigna la respuesta y se ejecuta la condicion
		fmt.Println(" La llave existe y su valor es", val)
	} else {
		fmt.Println("La llave no existe", val)
	}

	val1, ok1 := m1["robin"]

	if ok1 {
		fmt.Println(" La llave existe y su valor es", val1)
	} else {
		fmt.Println("La llave no existe", val1)
	}

	/*
		Si se trata de acceder a una llave qye no existe el resutlado es cero
	*/
}
