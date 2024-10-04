package Attendant

import (
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/ParkingLot"
	"parking-lot/Strategy"
	"parking-lot/Ticket"
)

type Attendant struct {
	assignedParkingLots []ParkingLot.ParkingLot
	parkingStrategy     Strategy.ParkingLotStrategy
}

func (attendant *Attendant) NewAttendant(strategy Strategy.ParkingLotStrategy) {
	attendant.assignedParkingLots = make([]ParkingLot.ParkingLot, 0)
	attendant.parkingStrategy = strategy
}

func (attendant *Attendant) assign(parkingLot ParkingLot.ParkingLot) error {
	for _, lot := range attendant.assignedParkingLots {
		if lot.IsSameParkingLot(parkingLot) {
			return customError.ErrParkingLotAlreadyAssigned
		}
	}
	attendant.assignedParkingLots = append(attendant.assignedParkingLots, parkingLot)
	return nil
}

func (attendant *Attendant) park(car *Car.Car) (Ticket.Ticket, error) {
	parkingLot, err := attendant.parkingStrategy.GetNextLot(attendant.assignedParkingLots)
	if err == nil {
		ticket, err := parkingLot.Park(car)
		return ticket, err
	}
	return Ticket.Ticket{}, customError.ErrParkingLotFull
}

func (attendant *Attendant) UnPark(ticket Ticket.Ticket) (Car.Car, error) {
	for _, parkingLot := range attendant.assignedParkingLots {
		car, err := parkingLot.UnPark(ticket)
		if err != nil {
			continue
		}
		return car, nil
	}
	return Car.Car{}, customError.ErrInvalidTicket
}
