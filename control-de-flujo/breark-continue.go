package main

import "fmt"

func main() {
	var x int = 1

	for {
		if x > 100 {
			break
		}

		if x%2 == 0 { // indica que el numero es par
			fmt.Println("Numero par ", x)
			x++
		} else {
			fmt.Println("Numero inpar ", x)
			x++
			continue // lo que hace es mandar a que itere
		}

	}
}
