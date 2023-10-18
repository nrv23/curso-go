package main

import (
	"fmt"
	"strconv"
)

type dinero2 int

var x int
var y dinero2
var cadena string = "123343"

func main() {

	y = 1
	x = int(y) //conversion de tipos. de tipo de dato dinero2 a int
	z, err := strconv.Atoi(cadena)

	if err != nil {
		fmt.Println("error: ", err)
	} else {

		fmt.Printf("el valor de y %v \n", y)
		fmt.Printf("el valor de x %v \n", x)
		fmt.Printf("el valor de cadena %v \n", z)

		fmt.Printf("el tipo de valor de y %T \n", y)
		fmt.Printf("el tipo de valor de x %T \n", x)

	}

}
