package strategy

import "parking-lot/parking_lot"

type ParkingLotStrategy interface {
	GetNextLot(parkingLots []parking_lot.ParkingLot) (parking_lot.ParkingLot, error)
}
