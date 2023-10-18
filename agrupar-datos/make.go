package main

import "fmt"

func main() {

	// literal compuesto
	x := make([]int, 10, 100)
	// el tercer paramtro de make es la capacidad del array sin re asignar memoria por parte de go
	// el array puede tener un tamano inicial de 10 pero crecer hasta 100, una vez que sse superan los 100 elementos go asigna memoria al array
	/*


		cuando hay reasignacion de memoria, go desecha el array original, asigna memoria dinamicamente para nuevo array y los carga con
		los elementos del anterior
	*/
	for i := 0; i < 10; i++ {
		x[i] = i
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1, 23, 324, 35, 435, 546, 54, 654, 6)
	x[17] = 100
	//capcidad del array
	fmt.Println(x)
	fmt.Println(len(x))
}
