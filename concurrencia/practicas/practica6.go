package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Arquitectura", runtime.GOOS)
	fmt.Println("Sistema Operativo", runtime.GOARCH)
}
