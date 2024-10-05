package ParkingLot

import (
	"fmt"
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/Slot"
	"parking-lot/ticket"
)

type ParkingLot struct {
	isFull       bool
	slots        []*Slot.Slot
	parkingLotId string
	ownedBy      string
}

func (parkingLot *ParkingLot) Construct(numberOfSlots int, owner string) error {
	if numberOfSlots < 1 {
		return customError.ErrSlotNumberShouldBeGreaterThanZero
	}
	parkingLot.isFull = false
	parkingLot.ownedBy = owner
	parkingLot.parkingLotId = fmt.Sprintf("%p", parkingLot)
	parkingLot.slots = make([]*Slot.Slot, numberOfSlots)
	for i := range parkingLot.slots {
		parkingLot.slots[i] = &Slot.Slot{}
	}
	return nil
}

func (parkingLot *ParkingLot) Park(car *Car.Car) (ticket.Ticket, error) {
	if car.IsCarParked() {
		return ticket.Ticket{}, customError.ErrCarAlreadyParked
	}
	for _, slot := range parkingLot.slots {
		if !slot.IsSlotOccupied() {
			ticket, _ := slot.Park(car)
			parkingLot.updateIsFull()
			return ticket, nil
		}
	}
	return ticket.Ticket{}, customError.ErrParkingLotFull
}

func (parkingLot *ParkingLot) UnPark(ticket ticket.Ticket) (Car.Car, error) {
	for _, slot := range parkingLot.slots {
		if slot.IsSlotOccupied() {
			car, err := slot.UnPark(ticket)
			if err != nil {
				continue
			}
			parkingLot.updateIsFull()
			return car, err
		}
	}
	return Car.Car{}, customError.ErrInvalidTicket
}

func (parkingLot *ParkingLot) CountCarsByColor(color Car.CarColor) int {
	count := 0
	for _, slot := range parkingLot.slots {
		if slot.IsSlotOccupied() && slot.IsCarColor(color) {
			count++
		}
	}
	return count
}

func (parkingLot *ParkingLot) GetCarParkedInfoByRegNo(registeredNumber string) (ticket.Ticket, error) {
	for _, slot := range parkingLot.slots {
		if slot.IsSlotOccupied() {
			ticket, err := slot.GetTicketIfCarMatches(registeredNumber)
			if err != nil {
				continue
			}
			return ticket, nil
		}
	}
	return ticket.Ticket{}, customError.ErrCarNotParked
}

func (parkingLot *ParkingLot) updateIsFull() {
	parkingLot.isFull = true
	for _, slot := range parkingLot.slots {
		if !slot.IsSlotOccupied() {
			parkingLot.isFull = false
			break
		}
	}
}

func (parkingLot *ParkingLot) IsSameParkingLot(receivedParkingLot ParkingLot) bool {
	return parkingLot.parkingLotId == receivedParkingLot.parkingLotId
}

func (parkingLot *ParkingLot) IsOwnedBy(id string) bool {
	return parkingLot.ownedBy == id
}

func (parkingLot *ParkingLot) IsParkingLotFull() bool {
	parkingLot.updateIsFull()
	return parkingLot.isFull
}

func (parkingLot *ParkingLot) GetAvailableSlots() int {
	count := 0
	for _, slot := range parkingLot.slots {
		if !slot.IsSlotOccupied() {
			count++
		}
	}
	return count
}
