package main

import (
	"errors"
)

var (
	ErrParkingLotAlreadyAssigned = errors.New("parking lot already assigned")
)

type Attendant struct {
	assignedParkingLots []ParkingLot
}

func (attendant *Attendant) assign(parkingLot ParkingLot) error {
	for i := 0; i < len(attendant.assignedParkingLots); i++ {
		if attendant.assignedParkingLots[i].isSameParkingLot(parkingLot) {
			return ErrParkingLotAlreadyAssigned
		}
	}
	attendant.assignedParkingLots = append(attendant.assignedParkingLots, parkingLot)
	return nil
}
