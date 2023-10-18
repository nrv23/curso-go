package main

import "fmt"

func main() {

	switch {

	case 2 == 2:
		{
			fmt.Println("Se imprime")
		}

	case 2 != 2:
		fmt.Print("No se imprime")
	default:
		fmt.Println("No se ejecuto ningun caso")
	}
}
