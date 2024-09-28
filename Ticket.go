package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ticket struct {
	ticketId string
}

func newTicket() Ticket {
	ticketId := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
	fmt.Println(ticketId)
	return Ticket{ticketId: ticketId}
}

func (ticket Ticket) validateTicket(receivedTicket Ticket) bool {
	return ticket.ticketId == receivedTicket.ticketId
}
