package main

import "fmt"

func main() {

	var deporteFav string = "opcion1"

	switch deporteFav {

	case "opcion1":
		fmt.Println(deporteFav)

	case "opcion2":
		fmt.Println("no es una opcion")

	default:
		fmt.Println("opcion default")
	}

}
