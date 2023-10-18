package main

import "fmt"

func main() {

	var x [5]int // declaro un arreglo entero de cinco lugares
	x[0] = 0
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Printf("%T \n", x)
	fmt.Println(len(x))
}
