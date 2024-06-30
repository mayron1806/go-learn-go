package main

import "fmt"

func printArrayData[T int | string](array []T) {
	// len = length
	// cap = capacity
	fmt.Println(array, len(array), cap(array))
	fmt.Println("-------------------------")
}
func main() {
	// slice
	names := []string{"Mayron", "Duda", "João"}
	printArrayData(names)
	names = append(names, "Sara")
	printArrayData(names)

	// cria slice
	// 1º parametro = tipo do slice
	// 2° tamanho do slice
	// 3º capacidade
	array := make([]int, 10, 20)
	printArrayData(array)
	fmt.Println("-------------------------")

	// map
	// não possui ordenação
	ages := make(map[string]int8)
	ages["Mayron"] = 20
	ages["Sara"] = 46
	ages["Gerssy"] = 50
	fmt.Println(ages)
	age, ok := ages["Joao"]
	fmt.Println(age, ok)

	fmt.Println("-------------------------")
	// struct
	type Person struct {
		Name   string
		Age    int8
		Status bool
		cpf    string // atrubuto privado
	}
	p := Person{
		Name:   "Mayron",
		Age:    20,
		Status: true,
		cpf:    "00000000000",
	}
	fmt.Println(p)

}
