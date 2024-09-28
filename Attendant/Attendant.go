package Attendant

import (
	"errors"
	"parking-lot/ParkingLot"
)

var (
	ErrParkingLotAlreadyAssigned = errors.New("parking lot already assigned")
)

type Attendant struct {
	assignedParkingLots []ParkingLot.ParkingLot
}

func (attendant *Attendant) assign(parkingLot ParkingLot.ParkingLot) error {
	for i := 0; i < len(attendant.assignedParkingLots); i++ {
		if attendant.assignedParkingLots[i].IsSameParkingLot(parkingLot) {
			return ErrParkingLotAlreadyAssigned
		}
	}
	attendant.assignedParkingLots = append(attendant.assignedParkingLots, parkingLot)
	return nil
}
