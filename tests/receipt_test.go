package tests

import (
	"parking-lot/receipt"
	"testing"
)

func TestTicketGenerationAndShouldValidateTrue(t *testing.T) {
	ticketOne := *receipt.Construct()

	if !ticketOne.ValidateTicket(ticketOne) {
		t.Error("Tickets should not be equal")
	}
}

func TestTwoTicketCannotBeSame(t *testing.T) {
	ticketOne := *receipt.Construct()
	ticketTwo := *receipt.Construct()

	if ticketOne.ValidateTicket(ticketTwo) {
		t.Error("Tickets should not be equal")
	}
}
