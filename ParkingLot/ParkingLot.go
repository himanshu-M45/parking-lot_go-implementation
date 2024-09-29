package ParkingLot

import (
	"fmt"
	"math/rand"
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/Slot"
	"parking-lot/Ticket"
	"time"
)

type ParkingLot struct {
	isFull       bool
	slots        []*Slot.Slot
	parkingLotId string
}

func (parkingLot *ParkingLot) NewParkingLot(numberOfSlots int) error {
	if numberOfSlots < 1 {
		return customError.ErrSlotNumberShouldBeGreaterThanZero
	}
	parkingLot.isFull = false
	parkingLot.slots = make([]*Slot.Slot, numberOfSlots)
	for i := range parkingLot.slots {
		parkingLot.slots[i] = &Slot.Slot{}
	}
	parkingLot.parkingLotId = fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
	return nil
}

func (parkingLot *ParkingLot) IsSameParkingLot(receivedParkingLot ParkingLot) bool {
	return parkingLot.parkingLotId == receivedParkingLot.parkingLotId
}

func (parkingLot *ParkingLot) IsParkingLotFull() bool {
	parkingLot.updateIsFull()
	return parkingLot.isFull
}

func (parkingLot *ParkingLot) Park(car *Car.Car) (Ticket.Ticket, error) {
	if car.IsCarParked() {
		return Ticket.Ticket{}, customError.ErrCarAlreadyParked
	}
	for _, slot := range parkingLot.slots {
		if !slot.IsSlotOccupied() {
			ticket, _ := slot.Park(car)
			parkingLot.updateIsFull()
			return ticket, nil
		}
	}
	return Ticket.Ticket{}, customError.ErrParkingLotFull
}

func (parkingLot *ParkingLot) UnPark(ticket Ticket.Ticket) (Car.Car, error) {
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

func (parkingLot *ParkingLot) GetCarParkedInfoByRegNo(registeredNumber string) (Ticket.Ticket, error) {
	for _, slot := range parkingLot.slots {
		if slot.IsSlotOccupied() {
			ticket, err := slot.GetTicketIfCarMatches(registeredNumber)
			if err != nil {
				continue
			}
			return ticket, nil
		}
	}
	return Ticket.Ticket{}, customError.ErrCarNotParked
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
