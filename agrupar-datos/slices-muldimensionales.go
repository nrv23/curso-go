package main

import "fmt"

func main() {

	x := []string{"nataniel", "venegas", "futbol", "pesas", "basket"}
	y := []string{"jose", "lopez", "futbol", "pesas", "juegos"}

	fmt.Println("")
	//fmt.Println(x, y)

	// crear un slice multidimensioonal

	vp := [][]string{x, y} // slice de dos dimensiones
	fmt.Println(vp)

}
