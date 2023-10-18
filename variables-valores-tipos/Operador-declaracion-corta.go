package main

import "fmt"

func main() {
	//...interface{} cantidad n de parametros de cualquier tipo
	fmt.Println("hola", 1, true)

	x := 23 // declaracion rapida de una variable con inferencia de tipo
	//var x int = 23 //declaracion con tipado estricto
	var y int = 23 + 2
	y = 1
	x = 2
	fmt.Println(x, y)
}
