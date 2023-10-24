package main

import "fmt"

func main() {

	// creacion de un canal
	/*
		Los canales envian y reciben valores
		chan es la palabra reservada de go para indicar que se refiere a un canal y luego el tipo de valor que envia y recibe el canal

	*/
	ca := make(chan int) // unbuffered channel o canal sin búfer

	// se crea la go rutina que recibe el valor del canal

	go func() {
		ca <- 42 // asignar valores a un canal
		/*
			La go rutina secundaria envia el valor por el canal y el programa no se bloquea
		*/

	}()
	fmt.Println(<-ca) /*
		La go rutina principal recibe el valor e imprime el valor del canal
	*/
}
