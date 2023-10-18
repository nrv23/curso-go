package main

import "fmt"

const (
	_ = iota // deseacha el 0, iota pasa a ser 1
	// actual  = 2023 iota aumenta su contador que pasa de 1 a 2

	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
	d = 2023 + iota
)

func main() {

	fmt.Println(a, b, c, d)
}
