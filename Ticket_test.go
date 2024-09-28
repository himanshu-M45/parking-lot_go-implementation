package main

import (
	"testing"
)

func TestTicketGenerationAndShouldValidateTrue(t *testing.T) {
	ticketOne := newTicket()

	if !ticketOne.validateTicket(ticketOne) {
		t.Error("Tickets should not be equal")
	}
}

func TestTwoTicketCannotBeSame(t *testing.T) {
	ticketOne := newTicket()
	ticketTwo := newTicket()

	if ticketOne.validateTicket(ticketTwo) {
		t.Error("Tickets should not be equal")
	}
}
