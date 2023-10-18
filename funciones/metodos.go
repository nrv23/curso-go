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

// metodo
// firma de una funcuin o metodo
// func (r receptor) identificador(parametros) (retorno(s)) {codigo}

func (a agentesecreto) presentarse() {
	/*

		los metodos son funciones o extensiones que se aplican al tipo de receptor del todo.
		funcion igual que las extensiones de .net
		con la exepcecion de que los metodos en go se extienden al tipo de receptor pero
		pueden tener ´parametros personalizados
	*/
	fmt.Println("Hola soy ", a.nombre+" ", a.apellido)
}

func main() {

	a1 := agentesecreto{
		persona: persona{
			nombre:   "nataniel",
			apellido: "vengeas",
		},

		lpm: true,
	}

	a2 := agentesecreto{
		persona: persona{
			nombre:   "nataniel",
			apellido: "rodriguez",
		},

		lpm: false,
	}

	fmt.Println(a1)
	a1.presentarse()
	a2.presentarse()
}
