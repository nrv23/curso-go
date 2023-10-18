package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	longitud float64
}

func (c circulo) area() float64 {
	return math.Pi * (c.radio * c.radio)
}

func (c cuadrado) area() float64 {
	return c.longitud * c.longitud
}

type IForma interface {
	area() float64
}

func Info(i IForma) {

	switch i.(type) {
	case circulo:
		fmt.Printf("imprimiendo el area desde circulo %v \n", i.(circulo).area())
		break
	case cuadrado:
		fmt.Printf("imprimiendo el area desde cuadrado %v \n", i.(cuadrado).area())

		break
	}
}

func main() {

	cir1 := circulo{
		radio: 100,
	}

	cua1 := cuadrado{
		longitud: 150,
	}

	Info(cir1)
	Info(cua1)
}
