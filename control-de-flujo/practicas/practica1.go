package main

import "fmt"

func main() {

	for i := 65; i < 91; i++ {
		fmt.Printf("%#U \n", i)
	}
}
