package main

import "fmt"

func main() {

	// declarar una variable de tipo Persona
	// estructuras anonimas que se declaran al mismo tiempo que la variable como una declaracion aninima
	// muy parecido a las interfaces aninomas de JS
	p1 := struct {
		persona struct {
			name string
			age  int
		}
		lpm bool
	}{
		lpm: true,
		persona: struct {
			name string
			age  int
		}{
			name: "nataniel",
			age:  31,
		},
	}

	fmt.Println(p1)
	fmt.Println(p1.persona)
	fmt.Println(p1.persona.age)
	fmt.Println(p1.persona.name)
	fmt.Println(p1.lpm)

}
