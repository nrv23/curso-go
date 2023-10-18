package main

import "fmt"

func main() {

	// literal compuesto
	x := []int{1, 2, 3, 4, 5, 6, 7, 8} // un array que se declara con sus valores

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	fmt.Println(x)
	fmt.Printf("%T", x)
}
