package main

import (
	"fmt"
	"log"

	m1 "github.com/mayron1806/go-learn-go/modules/module1"
)

func main() {
	var n1, n2 int
	_, err := fmt.Scan(&n1)
	if err != nil {
		log.Fatal("escreva um numero")
	}

	_, err = fmt.Scan(&n2)
	if err != nil {
		log.Fatal("escreva um numero")
	}

	res, err := m1.Soma(n1, n2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
