package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) // debe indicar el formato de cada variable para saber su tipo
	fmt.Println(s)
}
