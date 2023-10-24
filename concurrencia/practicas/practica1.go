package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() { // main es la go rutina principal

	wg.Add(2) // numero de go rutinas o funciones que se van a ejecutar al mismo tiempo.

	foo()
	go bar() // al indicarle la palabra reservada go a una funcion, se convierte automaticamente en una go rutina, un proceso
	// que se ejecuta de manera no secuencial en otro hilo distinto.
	go hi()
	fmt.Println(" CPUs \t", runtime.NumCPU()) // numero de nucleos disponibles
	fmt.Println(" go routines \t", runtime.NumGoroutine())

	wg.Wait() // indicar que espere si hay alguna go rutina sin ejecutar para que el main finalice
}

func foo() {
	fmt.Println("Se ejecuta primero que las go rutinas")
}

func bar() {
	fmt.Println("Go rutina 2")

	wg.Done() // le indica a go que la go rutina ya se ejecuto
	// cada vez que una go rutina se termina, el contador seteado con el metodo Add resta en 1.
}

func hi() {
	fmt.Println("Go rutina 3")

	wg.Done()
}
