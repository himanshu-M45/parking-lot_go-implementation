package Ticket

import (
	"testing"
)

func TestTicketGenerationAndShouldValidateTrue(t *testing.T) {
	ticketOne := *NewTicket()

	if !ticketOne.ValidateTicket(ticketOne) {
		t.Error("Tickets should not be equal")
	}
}

func TestTwoTicketCannotBeSame(t *testing.T) {
	ticketOne := *NewTicket()
	ticketTwo := *NewTicket()

	if ticketOne.ValidateTicket(ticketTwo) {
		t.Error("Tickets should not be equal")
	}
}
