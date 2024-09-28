package Ticket

import (
	"fmt"
	"math/rand"
	"time"
)

type Ticket struct {
	ticketId string
}

func NewTicket() *Ticket {
	ticketId := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
	return &Ticket{ticketId: ticketId}
}

func (ticket Ticket) ValidateTicket(receivedTicket Ticket) bool {
	return ticket.ticketId == receivedTicket.ticketId
}
