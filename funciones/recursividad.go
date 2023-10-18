package main

import "fmt"

func main() {

	fmt.Println(suma(100))
	fmt.Println(factorial(4))

}

func suma(x int) int {

	if x == 0 {
		return 1
	} else {
		x--
		return x + suma(x-1)
	}

}

func factorial(x int) int {

	if x == 0 {
		fmt.Println("llego", x)
		return 1
	} else {
		fmt.Println(x)
		return x * factorial(x-1)
	}
}
