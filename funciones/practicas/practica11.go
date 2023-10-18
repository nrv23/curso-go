package main

import "fmt"

func main() {
	fmt.Println(sumatoria(3))
}

func sumatoria(n int) int {

	if n == 1 {
		return 1
	} else {

		return n + sumatoria(n-1)
	}
}
