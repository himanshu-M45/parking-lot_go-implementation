package main

import (
	"testing"
	"time"
)

func TestTicketGenerationAndShouldValidateTrue(t *testing.T) {
	ticketOne := newTicket()

	if !ticketOne.validateTicket(ticketOne) {
		t.Error("Tickets should not be equal")
	}
}

func TestTwoTicketCannotBeSame(t *testing.T) {
	ticketOne := newTicket()
	time.Sleep(1 * time.Nanosecond)
	ticketTwo := newTicket()

	if ticketOne.validateTicket(ticketTwo) {
		t.Error("Tickets should not be equal")
	}
}
