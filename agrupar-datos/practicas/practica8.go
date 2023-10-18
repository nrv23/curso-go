package main

import "fmt"

func main() {

	x := map[string][]string{ // retorna en cada propiedad un array de tipo string
		`eduar_tua`:    {`computadoras`, `montañas`, `playa`}, // declarar que el valor es un array de tip
		`carlos_ramon`: {`leer`, `comprar`, `música`},
		`juan_bimba`:   {`helado`, `pintar`, `bailar`}, // entre llave se indica que es un array y se guardan los valores del arrauy
	}

	x["nataniel"] = []string{"computadora", "pesas", "futbol", "musica"}

	for key, value := range x {
		fmt.Println("registro ", key)

		if _, ok := x["nataniel"]; ok {
			delete(x, "nataniel")
		}

		for prop, el := range value {

			fmt.Println("Indice", prop)
			fmt.Println("valores", el)
		}
	}

	fmt.Println(x)
}
