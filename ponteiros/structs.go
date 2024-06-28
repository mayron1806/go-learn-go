package main

import (
	"github.com/google/uuid"
)

type Event struct {
	ID         string
	Name       string
	MaxTickets int
	Price      float32
	OwnerID    string
}
type Person struct {
	ID    string
	Name  string
	money float32
}

type Ticket struct {
	ID       string
	PersonID string
	EventID  string
	Price    float32
}

func buyTicket(e *Event, p *Person) (Ticket, error) {
	// verificar se tem dinheiro
	ticketPrice := e.Price

	if p.money < ticketPrice {
		return Ticket{}, PersonNoHaveMoney
	}
	// verificar se tem lugares
	if len(tickets) >= e.MaxTickets {
		return Ticket{}, AllTicketSelled
	}
	ticket := Ticket{
		ID:       uuid.NewString(),
		PersonID: p.ID,
		EventID:  e.ID,
		Price:    e.Price,
	}
	return ticket, nil
}
func createEvent(p *Person, name string, price float32, maxTickets int) (event Event) {
	event = Event{
		ID:         uuid.NewString(),
		Name:       name,
		MaxTickets: maxTickets,
		Price:      price,
		OwnerID:    p.ID,
	}
	return
}
