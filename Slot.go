package main

type Slot struct {
	car    *Car
	ticket *Ticket
}

func (slot Slot) newSlot() {
	slot.car = nil
	slot.ticket = nil
}
