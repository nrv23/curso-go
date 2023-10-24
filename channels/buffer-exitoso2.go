package main

import "fmt"

func main() {

	// creacion de un canal

	ca := make(chan int, 2) // buffered channel o canal con búfer
	/*
		el segunoo argumento del metodo make es el tamaño del buffer o cantidad de elementos del tipo de canal que
		van a ser almacenados.
	*/

	ca <- 42 // asignar valores a un canal
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
