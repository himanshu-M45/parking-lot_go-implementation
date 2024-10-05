package tests

import (
	"parking-lot/Car"
	"parking-lot/parking_lot"
	"parking-lot/role"
	"parking-lot/slots"
	"parking-lot/strategy"
)

var (
	owner    role.Owner
	newOwner role.Owner

	parkingLot  parking_lot.ParkingLot
	parkingLot1 parking_lot.ParkingLot
	parkingLot2 parking_lot.ParkingLot
	parkingLot3 parking_lot.ParkingLot
	parkingLot4 parking_lot.ParkingLot

	car  *Car.Car
	car1 *Car.Car
	car2 *Car.Car
	car3 *Car.Car

	attendant      role.Attendant
	newAttendant   role.Attendant
	smartAttendant role.Attendant

	slot slots.Slot
)

func Setup() {
	owner.Construct()
	newOwner.Construct()

	parkingLot, _ = owner.CreateParkingLot(1)
	parkingLot1, _ = owner.CreateParkingLot(1)
	parkingLot2, _ = owner.CreateParkingLot(2)
	parkingLot3, _ = newOwner.CreateParkingLot(1)
	parkingLot4, _ = newOwner.CreateParkingLot(3)

	attendant.Construct(&strategy.BasicLotStrategy{})
	newAttendant.Construct(&strategy.BasicLotStrategy{})
	smartAttendant.Construct(&strategy.SmartLotStrategy{})

	car = Car.Construct("KA-01-HH-1234", Car.BLACK)
	car1 = Car.Construct("KA-02-UY-9999", Car.WHITE)
	car2 = Car.Construct("KA-04-KL-6666", Car.BLACK)
	car3 = Car.Construct("KA-06-HJ-7777", Car.RED)

	slot.Construct()
}
