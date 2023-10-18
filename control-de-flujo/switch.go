package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write \n", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough // indica que se debe ejecutar el codigo del siguiente caso, en esta ocasion seria el 3
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { // una funcion con argumento de tipo cualquiera

		switch t := i.(type) { // saber el tipo de valor ingresado
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// es posible en el swtich tener en un caso varias expresiones y con una que sea verdadera el caso se ejecuta
	switch {
	case 2 == 2, 8 == 0, 9 == 9: // al menos una expresion debe ser verdadera
		fmt.Println("Caso se ejecuto")
	default:
		fmt.Println("No cayo en ningun caso")
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	/// operadores logicos en el switch

	switch {
	case 1 == 2 && 4 == 2:
		fmt.Println("caso 1")
	case 1 == 2 || 3 == 3:
		fmt.Println("caso 2")
	default:
		fmt.Println("No cayo en ningun caso")
	}
}
