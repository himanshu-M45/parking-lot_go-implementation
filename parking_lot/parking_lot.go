package parking_lot

import (
	"fmt"
	"parking-lot/Car"
	"parking-lot/common/custom_errors"
	"parking-lot/receipt"
	"parking-lot/slots"
)

type ParkingLot struct {
	isFull       bool
	slots        []*slots.Slot
	parkingLotId string
	ownedBy      string
}

func (parkingLot *ParkingLot) Construct(numberOfSlots int, owner string) error {
	if numberOfSlots < 1 {
		return custom_errors.ErrSlotNumberShouldBeGreaterThanZero
	}
	if owner == "" {
		return custom_errors.ErrCannotCreateParkingLotWithoutOwner
	}
	parkingLot.isFull = false
	parkingLot.ownedBy = owner
	parkingLot.parkingLotId = fmt.Sprintf("%p", parkingLot)
	parkingLot.slots = make([]*slots.Slot, numberOfSlots)
	for i := range parkingLot.slots {
		parkingLot.slots[i] = &slots.Slot{}
	}
	return nil
}

func (parkingLot *ParkingLot) Park(car *Car.Car) (receipt.Receipt, error) {
	if car.IsCarParked() {
		return receipt.Receipt{}, custom_errors.ErrCarAlreadyParked
	}
	for _, slot := range parkingLot.slots {
		if !slot.IsSlotOccupied() {
			ticket, _ := slot.Park(car)
			parkingLot.updateIsFull()
			return ticket, nil
		}
	}
	return receipt.Receipt{}, custom_errors.ErrParkingLotFull
}

func (parkingLot *ParkingLot) UnPark(ticket receipt.Receipt) (Car.Car, error) {
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
	return Car.Car{}, custom_errors.ErrInvalidTicket
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

func (parkingLot *ParkingLot) GetCarParkedInfoByRegNo(registeredNumber string) (receipt.Receipt, error) {
	for _, slot := range parkingLot.slots {
		if slot.IsSlotOccupied() {
			ticket, err := slot.GetTicketIfCarMatches(registeredNumber)
			if err != nil {
				continue
			}
			return ticket, nil
		}
	}
	return receipt.Receipt{}, custom_errors.ErrCarNotParked
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
