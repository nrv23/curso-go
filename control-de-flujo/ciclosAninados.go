package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("El ciclo externo %d \n ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("El ciclo interno %d \n", j)
		}
	}
}
