package test_setup

import (
	"parking-lot/ticket"
	"testing"
)

func TestTicketGenerationAndShouldValidateTrue(t *testing.T) {
	ticketOne := *ticket.Construct()

	if !ticketOne.ValidateTicket(ticketOne) {
		t.Error("Tickets should not be equal")
	}
}

func TestTwoTicketCannotBeSame(t *testing.T) {
	ticketOne := *ticket.Construct()
	ticketTwo := *ticket.Construct()

	if ticketOne.ValidateTicket(ticketTwo) {
		t.Error("Tickets should not be equal")
	}
}
