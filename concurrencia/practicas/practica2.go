package main

import "fmt"

type Persona struct { // funciona como una interfaz en JS
	name string
	age  int
}

func (p *Persona) hablar() {

	fmt.Println("Hola soy ", p.name+" ", p.age)
}

// definir una interfaz

type IHumano interface { // cualquier objeto que cumpla con la firma de esta  interfaz puede tiparse con esta interfaz
	hablar()
}

func dialogo(h IHumano) {
	h.hablar()
}

func main() {

	p1 := Persona{"Nataniel", 31}

	dialogo(&p1)

}
