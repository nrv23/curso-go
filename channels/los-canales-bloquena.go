package main

import "fmt"

func main() {

	// creacion de un canal
	/*
		Los canales envian y reciben valores
		chan es la palabra reservada de go para indicar que se refiere a un canal y luego el tipo de valor que envia y recibe el canal

	*/
	ca := make(chan int) // unbuffered channel o canal sin búfer
	ca <- 42             // asignar valores a un canal
	/*
		No hay un go rutine que reciba el valor del canal
		se bloquea el programa y se termina
	*/

	fmt.Println(<-ca)
}
