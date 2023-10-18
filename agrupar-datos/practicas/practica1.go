package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // un array que se declara con sus valores

	for _, val := range x {
		fmt.Println(val)
	}
}
