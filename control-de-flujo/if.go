package main

func main() {
	// se cumple solo si la condicion retorna true
	if true {
		println("Es un if verdadero")
	}

	if false {
		println("Es un if false")
	}

	if !false {
		println("Es un if !false")
	}

	if !true {
		println("Es un if !true")
	}
	if 42 == 42 {
		println("Es verdadero")
	}

	if x := 42; x == 42 { // asignacion y comparacion en elmismo if
		/*
			el scope de la variable x solo va existir en la condicion del if
		*/
		println("Es verdadero")
	}
}
