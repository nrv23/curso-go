package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(" OS \t", runtime.GOOS)
	fmt.Println(" AR \t", runtime.GOARCH)

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
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}

	wg.Done() // le indica a go que la go rutina ya se ejecuto
	// cada vez que una go rutina se termina, el contador seteado con el metodo Add resta en 1.
}

func hi() {
	for i := 0; i < 10; i++ {
		fmt.Println("hi:", i)
	}

	wg.Done()
}
