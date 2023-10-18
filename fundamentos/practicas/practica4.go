package main

import "fmt"

func main() {

	var a int = 4
	var b = a << 1 //hacer bigshoft a la derecha en una posicion

	fmt.Printf("%d \t %b \t %x \n", a, a, a)
	fmt.Printf("%d \t %b \t %x \n", b, b, b)
}
