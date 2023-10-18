package main

import "fmt"

const ( // declarar varias constantes en una sola declaracion

	a = iota
	b = iota
	c = iota
)

// el iota se resetea en cada declaracion unica de una constante

const ( // declarar varias constantes en una sola declaracion

	d = iota
	e = iota
	f = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)

	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", f)
}
