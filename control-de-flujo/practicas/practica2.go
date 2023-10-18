package main

import "fmt"

func main() {
	i := 1

	for i < 10001 {
		fmt.Println(i)
		i++
	}
}
