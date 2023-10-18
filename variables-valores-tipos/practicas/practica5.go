package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("EL tipo de la variable x es %T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("EL tipo de la variable y es %T\n", y)
}
