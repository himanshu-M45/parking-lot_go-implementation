package test_setup

import (
	"parking-lot/Attendant"
	"parking-lot/Car"
	"parking-lot/ParkingLot"
	"parking-lot/Slot"
	"parking-lot/Strategy"
)

var (
	owner    Attendant.Owner
	newOwner Attendant.Owner

	parkingLot  ParkingLot.ParkingLot
	parkingLot1 ParkingLot.ParkingLot
	parkingLot2 ParkingLot.ParkingLot
	parkingLot3 ParkingLot.ParkingLot
	parkingLot4 ParkingLot.ParkingLot

	car  *Car.Car
	car1 *Car.Car
	car2 *Car.Car
	car3 *Car.Car

	attendant      Attendant.Attendant
	newAttendant   Attendant.Attendant
	smartAttendant Attendant.Attendant

	slot Slot.Slot
)

func Setup() {
	owner.Construct()
	newOwner.Construct()

	parkingLot, _ = owner.CreateParkingLot(1)
	parkingLot1, _ = owner.CreateParkingLot(1)
	parkingLot2, _ = owner.CreateParkingLot(2)
	parkingLot3, _ = newOwner.CreateParkingLot(1)
	parkingLot4, _ = newOwner.CreateParkingLot(3)

	attendant.Construct(&Strategy.BasicLotStrategy{})
	newAttendant.Construct(&Strategy.BasicLotStrategy{})
	smartAttendant.Construct(&Strategy.SmartLotStrategy{})

	car = Car.Construct("KA-01-HH-1234", Car.BLACK)
	car1 = Car.Construct("KA-02-UY-9999", Car.WHITE)
	car2 = Car.Construct("KA-04-KL-6666", Car.BLACK)
	car3 = Car.Construct("KA-06-HJ-7777", Car.RED)

	slot.Construct()
}
