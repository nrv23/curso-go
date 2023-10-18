package main

import "fmt"

const cadena string = `Hola desde un raw literal no 
interpretado`

func main() {

	fmt.Printf("%s \n %T", cadena, cadena)
}
