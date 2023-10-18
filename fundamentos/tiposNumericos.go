package main

import (
	"fmt"
	"runtime"
)

var x uint16 // permite de 0 a 255 que no sea negativo
var y float32
var z uint8

func main() {
	x = 1000
	y = 1
	z = uint8(x)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)

	/// imprimir el sistema operitivo y la arquitectura de la maquina

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
