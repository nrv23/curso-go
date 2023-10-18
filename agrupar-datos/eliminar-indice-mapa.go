package main

import "fmt"

func main() {

	m1 := map[string]int{ //el nombre de la llave va en string y el valor en int
		"batman": 32,
		"robin":  40,
	}

	fmt.Println(m1)

	m1["nataniel"] = 31

	if v, ok := m1["robin"]; ok { // pregunta primero si existe el indice
		delete(m1, "robin") // borra el indice
		fmt.Println("Valor eliminado ", v)
	}

	if v, ok := m1["test"]; ok { // pregunta primero si existe el indice
		delete(m1, "test") // borra el indice
		fmt.Println("Valor eliminado ", v)
	} else {
		fmt.Println("El indice no existe")
	}
}
