package Attendant

import (
	"fmt"
	customError "parking-lot"
	"parking-lot/ParkingLot"
	"sync"
)

type Owner struct {
	ownerId          string
	ownedParkingLots []ParkingLot.ParkingLot
	parkingLotStatus sync.Map
	Attendant
}

func (owner *Owner) Construct() {
	owner.ownerId = fmt.Sprintf("%p", owner)
	owner.ownedParkingLots = make([]ParkingLot.ParkingLot, 0)
}

func (owner *Owner) CreateParkingLot(numberOfSlots int) (ParkingLot.ParkingLot, error) {
	parkingLot := ParkingLot.ParkingLot{}
	err := parkingLot.Construct(numberOfSlots, owner.ownerId)
	if err == nil {
		owner.ownedParkingLots = append(owner.ownedParkingLots, parkingLot)
		return parkingLot, nil
	}
	return ParkingLot.ParkingLot{}, err
}

func (owner *Owner) Assign(parkingLot ParkingLot.ParkingLot, attendant *Attendant) error {
	if owner.verifyParkingLot(parkingLot) {
		err := attendant.assign(parkingLot)
		if err == nil {
			return nil
		}
		return err
	}

	return customError.ErrOwnerDoesNotOwnParkingLot
}

func (owner *Owner) verifyParkingLot(parkingLot ParkingLot.ParkingLot) bool {
	for _, lot := range owner.ownedParkingLots {
		if lot.IsSameParkingLot(parkingLot) && lot.IsOwnedBy(owner.ownerId) {
			return true
		}
	}
	return false
}
