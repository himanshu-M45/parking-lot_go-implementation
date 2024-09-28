package Slot

import (
	"parking-lot/Car"
	"parking-lot/Ticket"
)

type Slot struct {
	car    *Car.Car
	ticket *Ticket.Ticket
}

func (slot Slot) NewSlot() {
	slot.car = nil
	slot.ticket = nil
}
