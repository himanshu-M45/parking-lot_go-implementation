package strategy

import (
	"parking-lot/common/custom_errors"
	"parking-lot/parking_lot"
)

type BasicLotStrategy struct{}

func (b *BasicLotStrategy) GetNextLot(parkingLots []parking_lot.ParkingLot) (parking_lot.ParkingLot, error) {
	for _, parkingLot := range parkingLots {
		if !parkingLot.IsParkingLotFull() {
			return parkingLot, nil
		}
	}
	return parking_lot.ParkingLot{}, custom_errors.ErrParkingLotFull
}
