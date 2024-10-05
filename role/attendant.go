package role

import (
	"parking-lot/Car"
	"parking-lot/common/custom_errors"
	"parking-lot/parking_lot"
	"parking-lot/receipt"
	"parking-lot/strategy"
)

type Attendant struct {
	assignedParkingLots []parking_lot.ParkingLot
	parkingStrategy     strategy.ParkingLotStrategy
}

func (attendant *Attendant) Construct(strategy strategy.ParkingLotStrategy) {
	attendant.assignedParkingLots = make([]parking_lot.ParkingLot, 0)
	attendant.parkingStrategy = strategy
}

func (attendant *Attendant) assign(parkingLot parking_lot.ParkingLot) error {
	for _, lot := range attendant.assignedParkingLots {
		if lot.IsSameParkingLot(parkingLot) {
			return custom_errors.ErrParkingLotAlreadyAssigned
		}
	}
	attendant.assignedParkingLots = append(attendant.assignedParkingLots, parkingLot)
	return nil
}

func (attendant *Attendant) Park(car *Car.Car) (receipt.Receipt, error) {
	parkingLot, err := attendant.parkingStrategy.GetNextLot(attendant.assignedParkingLots)
	if err == nil {
		ticket, err := parkingLot.Park(car)
		return ticket, err
	}
	return receipt.Receipt{}, custom_errors.ErrParkingLotFull
}

func (attendant *Attendant) UnPark(ticket receipt.Receipt) (Car.Car, error) {
	for _, parkingLot := range attendant.assignedParkingLots {
		car, err := parkingLot.UnPark(ticket)
		if err != nil {
			continue
		}
		return car, nil
	}
	return Car.Car{}, custom_errors.ErrInvalidTicket
}
