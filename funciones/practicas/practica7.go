package main

import "fmt"

func main() {

	suma := func(num1 int, num2 int) int {

		return num1 + num2
	}

	fmt.Println("La suma es ", suma(10, 20))

}
