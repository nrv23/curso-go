package main

import "fmt"

func main() {

	x := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[2:7]) //n impresos: [42 43 44 45 46
	fmt.Println(x[7:])  //n impresos: [47 48 49 50 51]
	fmt.Println(x[4:9]) //n impresos: [44 45 46 47 48]
	fmt.Println(x[3:8]) //n impresos: [43 44 45 46 47]

}
