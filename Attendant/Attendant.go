package Attendant

import (
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/ParkingLot"
	"parking-lot/Ticket"
)

type Attendant struct {
	assignedParkingLots []ParkingLot.ParkingLot
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
	for _, parkingLot := range attendant.assignedParkingLots {
		if !parkingLot.IsParkingLotFull() {
			return parkingLot.Park(car)
		}
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
