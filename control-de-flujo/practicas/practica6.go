package main

import "fmt"

func main() {

	// los valores declarados no tienen valor por dwfault entonces toma iota, a = 0, b =1
	var a int
	fmt.Println("Ingrese el primer valor")
	fmt.Scanf("%d\n", &a) // leer datos desde la consola

	var b int
	fmt.Println("Ingrese el segundo valor")
	fmt.Scanf("%d\n", &b) // leer datos desde la consola

	if a > b {
		fmt.Printf("el valor mayor es %v \n", a)
	} else if b > a {
		fmt.Printf("el valor mayor es %v\n", b)
	} else {
		fmt.Println("valores son iguales")
	}
}
