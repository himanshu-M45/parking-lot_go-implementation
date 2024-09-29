package Slot

import (
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/Ticket"
)

type Slot struct {
	car    *Car.Car
	ticket *Ticket.Ticket
}

func (slot *Slot) NewSlot() {
	slot.car = nil
	slot.ticket = nil
}

func (slot *Slot) IsSlotOccupied() bool {
	return slot.car != nil
}

func (slot *Slot) Park(car *Car.Car) (Ticket.Ticket, error) {
	if slot.car != nil {
		return Ticket.Ticket{}, customError.ErrCarAlreadyParked
	}
	slot.ticket = Ticket.NewTicket()
	slot.car = car
	slot.car.SetCarParked(true)
	return *slot.ticket, nil
}

func (slot *Slot) IsCarColor(color Car.CarColor) bool {
	return slot.car.IsSameColor(color)
}
