package slots

import (
	"parking-lot/Car"
	"parking-lot/common/custom_errors"
	"parking-lot/receipt"
)

type Slot struct {
	car    *Car.Car
	ticket *receipt.Receipt
}

func (slot *Slot) Construct() {
	slot.car = nil
	slot.ticket = nil
}

func (slot *Slot) IsSlotOccupied() bool {
	return slot.car != nil
}

func (slot *Slot) Park(car *Car.Car) (receipt.Receipt, error) {
	if slot.car != nil {
		return receipt.Receipt{}, custom_errors.ErrCarAlreadyParked
	}
	slot.ticket = receipt.Construct()
	slot.car = car
	slot.car.SetCarParked(true)
	return *slot.ticket, nil
}

func (slot *Slot) UnPark(ticket receipt.Receipt) (Car.Car, error) {
	if slot.car == nil || !slot.ticket.ValidateTicket(ticket) {
		return Car.Car{}, custom_errors.ErrInvalidTicket
	}
	slot.car.SetCarParked(false)
	car := *slot.car
	slot.car = nil
	slot.ticket = nil
	return car, nil
}

func (slot *Slot) IsCarColor(color Car.CarColor) bool {
	return slot.car.IsSameColor(color)
}

func (slot *Slot) GetTicketIfCarMatches(registeredNumber string) (receipt.Receipt, error) {
	if slot.car != nil && slot.car.IsIdenticalCar(registeredNumber) {
		return *slot.ticket, nil
	}
	return receipt.Receipt{}, custom_errors.ErrCarNotParked
}
