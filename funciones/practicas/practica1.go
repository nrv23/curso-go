package main

func main() {

	println("El valor escogido es ", foo(1))

	num, s := bar(1, "hola")
	println("Valores escogidos ", num, s)

}

func foo(num int) int {

	return num
}

func bar(num int, s string) (int, string) {

	return num, s
}
