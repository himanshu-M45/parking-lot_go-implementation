package strategy

import (
	customError "parking-lot"
	"parking-lot/parking_lot"
)

type SmartLotStrategy struct{}

func (s *SmartLotStrategy) GetNextLot(parkingLots []parking_lot.ParkingLot) (parking_lot.ParkingLot, error) {
	bestParkingLot := parking_lot.ParkingLot{}
	maxAvailableSlots := 0
	for _, parkingLot := range parkingLots {
		availableSlots := parkingLot.GetAvailableSlots()
		if availableSlots > maxAvailableSlots {
			maxAvailableSlots = availableSlots
			bestParkingLot = parkingLot
		}
	}
	if maxAvailableSlots > 0 && !bestParkingLot.IsParkingLotFull() {
		return bestParkingLot, nil
	}
	return parking_lot.ParkingLot{}, customError.ErrParkingLotFull
}
