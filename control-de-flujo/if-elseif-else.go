package main

func main() {
	// se cumple solo si la condicion retorna true

	const (
		x = 10
		y = 2
		z = 30
	)

	if x < y {
		println(x)
	} else if x > z {
		println(z)
	} else {
		println(y)
	}
}
