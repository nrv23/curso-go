package main

func main() {

	println(foo(1, 2, 3, 4, 4, 45, 5, 56, 56, 57, 6570))
	println(bar([]int{1, 3, 3443, 55, 6, 577, 8, 789, 980}))

}

func foo(a ...int) int {

	var suma int = 0
	for i := 0; i < len(a); i++ {
		suma += a[i]
	}

	return suma
}

func bar(a []int) int {
	var suma int = 0
	for i := 0; i < len(a); i++ {
		suma += a[i]
	}

	return suma
}
