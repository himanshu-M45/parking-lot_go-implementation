package Strategy

import "parking-lot/ParkingLot"

type ParkingLotStrategy interface {
	GetNextLot(parkingLots []ParkingLot.ParkingLot) (ParkingLot.ParkingLot, error)
}
