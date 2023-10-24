package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//runtime.GOMAXPROCS(2)
	fmt.Println(" CPUs \t", runtime.NumCPU()) // numero de nucleos disponibles
	fmt.Println(" go routines \t", runtime.NumGoroutine())

	var contador int64 = 0
	const gs int = 100
	var wg sync.WaitGroup
	var mtx sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock() // este metodo lock bloquea el acceso a la variable mientras un go rutine este leyendo o escribiendo sobre la variable
			v := contador
			v++
			//atomic.AddInt64(&contador, 1)
			//runtime.Gosched() // esta funcion cede el hilo a otra go rutina pero no suspende la gorutina actual. Esto se conoce como yield

			contador = v
			mtx.Unlock() // desbloquear el acceso para que otra go rutine puede acceder a la variable
			fmt.Println("contador ", atomic.LoadInt64(&contador))
			wg.Done()
		}()

		fmt.Println(" go routines restantes \t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("cuenta ", contador)
}
