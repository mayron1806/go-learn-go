package main

import "fmt"

type Person struct {
	Name   string
	Status bool
}

func (p Person) String() string {
	return fmt.Sprintf("%s teste", p.Name)
}

type FisicPerson struct {
	Person
	CPF string
}
type JuridicPerson struct {
	CNPJ string
}

func main() {
	p := FisicPerson{
		Person: Person{Name: "Mayron", Status: true},
		CPF:    "00000000000",
	}
	fmt.Println(p)
}
