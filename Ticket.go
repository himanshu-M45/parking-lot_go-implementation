package main

import "time"

type Ticket struct {
	ticketId int64
}

func newTicket() Ticket {
	return Ticket{ticketId: time.Now().UnixNano()}
}

func (ticket Ticket) validateTicket(receivedTicket Ticket) bool {
	return ticket.ticketId == receivedTicket.ticketId
}
