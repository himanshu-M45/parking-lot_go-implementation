package role

import (
	"fmt"
	"parking-lot/common/custom_errors"
	"parking-lot/parking_lot"
	"parking-lot/strategy"
	"sync"
)

type Owner struct {
	ownerId          string
	ownedParkingLots []parking_lot.ParkingLot
	parkingLotStatus sync.Map
	*Attendant
}

func (owner *Owner) Construct() {
	owner.ownerId = fmt.Sprintf("%p", owner)
	owner.ownedParkingLots = make([]parking_lot.ParkingLot, 0)
	owner.Attendant = &Attendant{}
	owner.Attendant.Construct(&strategy.SmartLotStrategy{})
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

	return custom_errors.ErrOwnerDoesNotOwnParkingLot
}

func (owner *Owner) verifyParkingLot(parkingLot parking_lot.ParkingLot) bool {
	for _, lot := range owner.ownedParkingLots {
		if lot.IsSameParkingLot(parkingLot) && lot.IsOwnedBy(owner.ownerId) {
			return true
		}
	}
	return false
}
