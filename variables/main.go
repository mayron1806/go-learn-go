package main

import (
	"fmt"
	"math"
)

func main() {
	// variaveis
	fmt.Println("Variavies ------------------------")
	var boolean bool
	fmt.Println(boolean)

	str := "ola mundo"
	fmt.Println(str)

	var int1, int2 int = 1, 2
	fmt.Println(int1, int2)

	f1, f2 := 1, 2
	fmt.Println(f1, f2)

	// constantes
	fmt.Println("Constantes ------------------------")
	const pi = math.Pi
	fmt.Println(pi)
}
