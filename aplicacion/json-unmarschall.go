package main

import (
	"encoding/json"
	"fmt"
)

/*
type Persona []struct { // la estructura indica que es un array de tipo Persona

		Nombre   string `json:"Nombre"`
		Apellido string `json:"Apellido"`
		Edad     int    `json:"Edad"`
	}
*/
type Persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {

	var jsonBytes = []byte(`[{"Nombre":"Nataniel","Apellido":"Vengeas","Edad":32},{"Nombre":"Nataniel","Apellido":"Rodríguez","Edad":31}]`)
	p1 := []Persona{} // se debe pasar un puntero de una estructura para que Unmarshal pueda mapear y convertir el json con el valor del tipo de estructura
	// y asi devolver el json de retorno
	err := json.Unmarshal(jsonBytes, &p1) // devuelve el json parseado para leerlo o manipularlo
	/*
		una vez procesado el puntero por unmarschally que no devuelva error, el valor que esta dentro del puntero guarda el objeto json,
		que es el valor que se puede leer o manipular
	*/

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("json generado", p1)

	// se puede recorrer el json

	for _, elemento := range p1 {
		fmt.Println("Nombre", elemento.Nombre)
		fmt.Println("Apellido", elemento.Apellido)
		fmt.Println("Edad", elemento.Edad)
	}

	p1[0].Nombre = "Nombre cambiado"
	p1[0].Edad = 18

	fmt.Println("json generado", p1)
}
