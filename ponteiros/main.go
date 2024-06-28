package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

var PersonNoHaveMoney = errors.New("you don`t have money to buy this ticket for this event")
var AllTicketSelled = errors.New("tickets allready selled")
var tickets []Ticket

func main() {
	start := time.Now().UnixMilli()
	fake := faker.New()
	var persons []Person
	nPersons := rand.IntN(1000) + 1
	for i := 0; i < nPersons; i++ {
		persons = append(persons, Person{
			ID:    uuid.NewString(),
			money: rand.Float32() * 10000,
			Name:  fake.Person().FirstName(),
		})
	}
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i])
	}
	end := time.Now().UnixMilli()

	duration := end - start
	fmt.Println(duration)
}
