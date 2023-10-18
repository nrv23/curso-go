package main

import "fmt"

const (
	a     = 34
	b int = 23
)

const (
	c string = "Hola mundo"
	d        = iota
)

func main() {

	fmt.Printf("%v \t\t %T \n", a, a)

	fmt.Printf("%v \t\t %T \n", b, b)
	fmt.Printf("%v \t\t %T \n", c, c)
	fmt.Printf("%v \t\t %T \n", d, d)
}
