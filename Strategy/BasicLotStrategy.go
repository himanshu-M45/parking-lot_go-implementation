package Strategy

import (
	customError "parking-lot"
	"parking-lot/ParkingLot"
)

type BasicLotStrategy struct{}

func (b *BasicLotStrategy) GetNextLot(parkingLots []ParkingLot.ParkingLot) (ParkingLot.ParkingLot, error) {
	for _, parkingLot := range parkingLots {
		if !parkingLot.IsParkingLotFull() {
			return parkingLot, nil
		}
	}
	return ParkingLot.ParkingLot{}, customError.ErrParkingLotFull
}
