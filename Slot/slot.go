package Slot

import (
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/ticket"
)

type Slot struct {
	car    *Car.Car
	ticket *ticket.Ticket
}

func (slot *Slot) Construct() {
	slot.car = nil
	slot.ticket = nil
}

func (slot *Slot) IsSlotOccupied() bool {
	return slot.car != nil
}

func (slot *Slot) Park(car *Car.Car) (ticket.Ticket, error) {
	if slot.car != nil {
		return ticket.Ticket{}, customError.ErrCarAlreadyParked
	}
	slot.ticket = ticket.Construct()
	slot.car = car
	slot.car.SetCarParked(true)
	return *slot.ticket, nil
}

func (slot *Slot) UnPark(ticket ticket.Ticket) (Car.Car, error) {
	if slot.car == nil || !slot.ticket.ValidateTicket(ticket) {
		return Car.Car{}, customError.ErrInvalidTicket
	}
	car := *slot.car
	slot.car.SetCarParked(false)
	slot.car = nil
	slot.ticket = nil
	return car, nil
}

func (slot *Slot) IsCarColor(color Car.CarColor) bool {
	return slot.car.IsSameColor(color)
}

func (slot *Slot) GetTicketIfCarMatches(registeredNumber string) (ticket.Ticket, error) {
	if slot.car != nil && slot.car.IsIdenticalCar(registeredNumber) {
		return *slot.ticket, nil
	}
	return ticket.Ticket{}, customError.ErrCarNotParked
}
