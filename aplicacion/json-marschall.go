package main

import (
	"encoding/json" // usar esta libreria para los datos de tipo json, ya sea parsear o convertir a string
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := Persona{
		"Nataniel", "Vengeas", 32,
	}

	p2 := Persona{
		"Nataniel", "Rodríguez", 31,
	}

	personas := []Persona{p1, p2}
	bs, err := json.Marshal(personas)
	// esto devuelve un array de bytes o un error, el array de bytes
	// el array de bytes se debe convertir a string

	if err != nil {
		fmt.Println("Hubo un error", err)
		return
	}
	jsonOutput := string(bs) // una vez convertido en un string se puede pintar en pantalla
	fmt.Println(jsonOutput)
}
