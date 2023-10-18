/*
	usar estas expresiones
		== igual
		<= mayor o igual
		>= menor o igual
		!= diferente
		< mayor que
		> menor que

*/

package main

import "fmt"

func main() {

	var (
		a = 10
		b = 20
		c = 30
		d = 10
	)

	fmt.Println(a > b)
	fmt.Println(a == d)
	fmt.Println(b <= c)
	fmt.Println(c >= d)
	fmt.Println(c > b)
	fmt.Println(b != a)

}
