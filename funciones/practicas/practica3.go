package main

func main() {
	defer bar()
	foo()

}

func foo() {

	println("Hola...")
}

func bar() {

	println("Adios...")
}
