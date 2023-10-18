package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}

	u2 := usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	u3 := usuario{
		Nombre: "Che",
		Edad:   54,
	}

	usuarios := []usuario{u1, u2, u3}

	bs, err := json.Marshal(usuarios)
	// esto devuelve un array de bytes o un error, el array de bytes
	// el array de bytes se debe convertir a string

	if err != nil {
		fmt.Println("Hubo un error", err)
		return
	}
	jsonOutput := string(bs) // una vez convertido en un string se puede pintar en pantalla
	fmt.Println(jsonOutput)

	//tu código va aquí
}
