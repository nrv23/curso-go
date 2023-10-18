package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s1 string = "Hola mundo" // string literal interpretado
	var s2 string = `Esta es una cadena 
	partida` // row string literal no interpretado

	fmt.Println(s1)
	fmt.Printf("%T \n", s1)

	fmt.Println(s2)
	fmt.Printf("%T \n", s2)

	// declarar un slice de bytes y luego convertir un string a un array de bytes
	// al convertir un string a un array de bytes, cada letra del string se convierte al codigo ascii que le pertenece
	// por eso el resultado seria []uint8 , son numeros del 0 - 255 no negativos

	fmt.Println("")
	bs := []byte(s1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)

	for i := 0; i < len(s1); i++ {
		b := byte(s1[i])
		fmt.Println("")
		fmt.Printf("%v \n", b)      // codigo ascii de el caracter
		fmt.Printf("%#U \n", s1[i]) // codigo unicode
	}

	// range o foreach
	fmt.Println("")
	fmt.Println("")

	for i, v := range s1 { // v es la iteracion por cada elemento del array
		b := byte(v)
		c := strconv.Itoa(int(v))

		fmt.Printf("En el indice %d  el valor es %v \n", i, b)
		fmt.Printf("En el indice %d  el valor es %q \n", i, b)
		fmt.Printf("En el indice %d  el valor es %x \n", i, b)
		fmt.Printf("En el indice %d  el valor es %s \n", i, c)

		//fmt.Printf("%v \n", b)  // codigo ascii de el caracter
		fmt.Printf("%#U \n", v) // codigo unicode

	}
}
