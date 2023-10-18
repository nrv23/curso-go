package main

import "fmt"

func main() {

	x := [5]int{} // un array que se declara con sus valores

	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	for _, val := range x {
		fmt.Println(val)
	}
}
