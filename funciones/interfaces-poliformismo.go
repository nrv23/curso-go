package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agentesecreto struct {
	persona
	lpm bool
}

func (a agentesecreto) presentarse() {

	fmt.Println("Hola soy ", a.nombre+" ", a.apellido, "EL agente secreto se presentó")
}

func (p persona) presentarse() {

	fmt.Println("Hola soy ", p.nombre+" ", p.apellido)
}

// definir una interfaz

type IHumano interface { // cualquier objeto que cumpla con la firma de esta  interfaz puede tiparse con esta interfaz
	presentarse()
}

func recibeHumano(h IHumano) {

	switch h.(type) { // h.(type) obtener el tipo de dato de la variable de entrada h. Funciona como un typeof en Javascript
	case persona:
		// acceder a los valores de persona
		fmt.Println("imprimiendo el nombre desde persona", h.(persona).nombre)
		break
	case agentesecreto:
		// acceder a los valores de agentesecreto
		fmt.Println("imprimiendo el nombre desde agente", h.(agentesecreto).nombre)
		/*
			h.(agentesecreto) esta 3expresion es una asercion. Le estoy diciendo que h es de tipo agentesecreto y puedo accerder a los valores
			o funciones que tenga ese tipo

		*/
		break
	default:
		h.presentarse()
	}
}

type manzana int

func main() {

	a1 := agentesecreto{
		persona: persona{
			nombre:   "nataniel",
			apellido: "vengeas",
		},

		lpm: true,
	}

	p1 := persona{
		nombre:   "nataniel2",
		apellido: "rodriguez",
	}

	fmt.Println(a1)
	a1.presentarse()
	p1.presentarse()
	fmt.Println("Imprimido desde la funcion que tiene el parametro de interfaz")
	recibeHumano(a1)
	recibeHumano(p1)

	var x manzana = 42
	var y int = int(x) // conversion explicita

	fmt.Println(y)

}
