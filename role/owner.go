package role

import (
	"fmt"
	customError "parking-lot"
	"parking-lot/parking_lot"
	"sync"
)

type Owner struct {
	ownerId          string
	ownedParkingLots []parking_lot.ParkingLot
	parkingLotStatus sync.Map
	Attendant
}

func (owner *Owner) Construct() {
	owner.ownerId = fmt.Sprintf("%p", owner)
	owner.ownedParkingLots = make([]parking_lot.ParkingLot, 0)
}

func (owner *Owner) CreateParkingLot(numberOfSlots int) (parking_lot.ParkingLot, error) {
	parkingLot := parking_lot.ParkingLot{}
	err := parkingLot.Construct(numberOfSlots, owner.ownerId)
	if err == nil {
		owner.ownedParkingLots = append(owner.ownedParkingLots, parkingLot)
		return parkingLot, nil
	}
	return parking_lot.ParkingLot{}, err
}

func (owner *Owner) Assign(parkingLot parking_lot.ParkingLot, attendant *Attendant) error {
	if owner.verifyParkingLot(parkingLot) {
		err := attendant.assign(parkingLot)
		if err == nil {
			return nil
		}
		return err
	}

	return customError.ErrOwnerDoesNotOwnParkingLot
}

func (owner *Owner) verifyParkingLot(parkingLot parking_lot.ParkingLot) bool {
	for _, lot := range owner.ownedParkingLots {
		if lot.IsSameParkingLot(parkingLot) && lot.IsOwnedBy(owner.ownerId) {
			return true
		}
	}
	return false
}
