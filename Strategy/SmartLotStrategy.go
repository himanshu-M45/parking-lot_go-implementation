package Strategy

import (
	customError "parking-lot"
	"parking-lot/ParkingLot"
)

type SmartLotStrategy struct{}

func (s *SmartLotStrategy) GetNextLot(parkingLots []ParkingLot.ParkingLot) (ParkingLot.ParkingLot, error) {
	bestParkingLot := ParkingLot.ParkingLot{}
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
	return ParkingLot.ParkingLot{}, customError.ErrParkingLotFull
}
