package main

import "fmt"

func main() {

	defer foo() // se espera hasta que la fucnion que la contiene ejecute todo el codigo y cuando llega al return o final ejecuta
	// la funcion cin defer
	defer hello()
	bar()

}

func foo() {

	fmt.Println("foo")
}

func bar() {

	fmt.Println("bar")
}

func hello() {

	fmt.Println("hello")
}

/*

	Lo que hace defer es que pospne la ejecucion de la funcion hasta que la funcion que la envuelve retorne o llegue al final

*/
