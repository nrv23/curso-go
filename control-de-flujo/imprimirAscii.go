package main

import "fmt"

func main() {
	for i := 31; i < 130; i++ {
		fmt.Printf("%#U \n", i)
	}
}
