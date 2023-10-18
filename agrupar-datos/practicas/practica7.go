package main

import "fmt"

func main() {
	/*
		"Batman", "Jefe", "Vestido de negro"
		"Robin", "Ayudante", "Vestido de colores"
	*/

	vp := [][]string{}

	p1 := []string{"Batman", "Jefe", "Vestido de negro"}
	p2 := []string{"Robin", "Ayudante", "Vestido de colores"}

	vp = append(vp, p1, p2)

	for key, reg := range vp {
		fmt.Println("Registro:", key)
		for j, val := range reg { // el segundo for muestra los arrays internos
			fmt.Printf("\tÍndice: %v\tvalor: %v\n", j, val)
		}
	}
}
