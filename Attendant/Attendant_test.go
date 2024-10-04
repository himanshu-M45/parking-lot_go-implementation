package Attendant

import (
	"errors"
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/ParkingLot"
	"parking-lot/Strategy"
	"testing"
)

var (
	attendant  Attendant
	parkingLot ParkingLot.ParkingLot
	car        *Car.Car
)

func setup() {
	attendant = Attendant{}
	attendant.NewAttendant(&Strategy.BasicLotStrategy{})
	parkingLot = ParkingLot.ParkingLot{}
	_ = parkingLot.NewParkingLot(2)
	_ = attendant.assign(parkingLot)
	car = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
}

// ------------------------------- assign parkingLot tests -------------------------------
func TestAssignParkingLotToAttendant(t *testing.T) {
	setup()
	newParkingLot := ParkingLot.ParkingLot{}
	_ = newParkingLot.NewParkingLot(1)
	err := attendant.assign(newParkingLot)
	if err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAssignTwoParkingLotToAttendant(t *testing.T) {
	setup()
	parkingLot2 := ParkingLot.ParkingLot{}
	_ = parkingLot2.NewParkingLot(1)
	if err := attendant.assign(parkingLot2); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCannotAssignSameParkingLotTwice(t *testing.T) {
	setup()
	err := attendant.assign(parkingLot)
	if !errors.Is(err, customError.ErrParkingLotAlreadyAssigned) { // !True
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotAlreadyAssigned, err)
	}
}

func TestAssignSameParkingLotToMultipleAttendants(t *testing.T) {
	setup()
	attendant2 := Attendant{}
	if err := attendant2.assign(parkingLot); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

// ------------------------------- park through attendant tests -------------------------------
func TestParkingCarThroughAttendant(t *testing.T) {
	setup()

	_, err := attendant.park(car)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestTryToParkCarWhenParkingLotIsFull(t *testing.T) {
	setup()
	secondCar, thirdCar := &Car.Car{}, &Car.Car{}
	secondCar, thirdCar = Car.NewCar("KA-01-HH-9999", Car.GREEN), Car.NewCar("KA-01-HH-1234", Car.BLUE)

	_, _ = attendant.park(car)
	_, _ = attendant.park(secondCar)
	_, err := attendant.park(thirdCar)

	if !errors.Is(err, customError.ErrParkingLotFull) {
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotFull, err)
	}
}

func TestParkSameCarTwice(t *testing.T) {
	setup()

	_, _ = attendant.park(car)
	_, err := attendant.park(car)

	if !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
}

func TestParkMultipleCarsInParkingLotThroughAttendant(t *testing.T) {
	setup()
	secondCar := &Car.Car{}
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant.park(car)
	_, err := attendant.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkCarInMultipleParkingLotsThroughSameAttendant(t *testing.T) {
	setup()
	parkingLot2, secondCar := ParkingLot.ParkingLot{}, &Car.Car{}
	_, _ = parkingLot2.NewParkingLot(1), attendant.assign(parkingLot2)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant.park(car)
	_, err := attendant.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkSameCarInDifferentParkingLotsOfSameAttendant(t *testing.T) {
	setup()
	parkingLot2 := ParkingLot.ParkingLot{}
	_, _ = parkingLot2.NewParkingLot(1), attendant.assign(parkingLot2)

	_, _ = attendant.park(car)
	_, err := attendant.park(car)

	if !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
}

func TestParkCarInDifferentParkingLotThroughDifferentAttendants(t *testing.T) {
	setup()
	parkingLot1, parkingLot2, attendant1, attendant2 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}, Attendant{}
	secondCar := &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1)
	attendant1.NewAttendant(&Strategy.BasicLotStrategy{})
	attendant2.NewAttendant(&Strategy.BasicLotStrategy{})
	_, _ = attendant1.assign(parkingLot1), attendant2.assign(parkingLot2)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant1.park(car)
	_, err := attendant2.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkCarInSameParkingLotThroughDifferentAttendants(t *testing.T) {
	setup()
	attendant2, secondCar := Attendant{}, &Car.Car{}
	attendant2.NewAttendant(&Strategy.BasicLotStrategy{})
	_, secondCar = attendant2.assign(parkingLot), Car.NewCar("KA-01-HH-9999", Car.BLUE)

	_, _ = attendant.park(car)
	_, err := attendant2.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// ------------------------------- unpark through attendant Tests -------------------------------
func TestUnParkCarThroughAttendant(t *testing.T) {
	setup()
	ticket, _ := attendant.park(car)
	receivedCar, _ := attendant.UnPark(ticket)

	if !receivedCar.IsIdenticalCar("KA-01-HH-1234") {
		t.Errorf("Expected unpaked car")
	}
}

func TestUnParkCarFromMultipleParkingLots(t *testing.T) {
	setup()
	parkingLot1, parkingLot2, attendant1, secondCar := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}
	attendant1.NewAttendant(&Strategy.BasicLotStrategy{})
	_, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1)
	_, _ = attendant1.assign(parkingLot1), attendant1.assign(parkingLot2)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	firstCarTicket, _ := attendant1.park(car)
	secondCarTicket, _ := attendant1.park(secondCar)

	_, _ = attendant1.UnPark(firstCarTicket)
	_, err := attendant1.UnPark(secondCarTicket)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUnParkAlreadyUnParkedCar(t *testing.T) {
	setup()
	ticket, _ := attendant.park(car)
	_, _ = attendant.UnPark(ticket)
	_, err := attendant.UnPark(ticket)

	if !errors.Is(err, customError.ErrInvalidTicket) {
		t.Errorf("Expected error '%v', got %v", customError.ErrInvalidTicket, err)
	}
}

func TestUnParkMultipleCarsFromMultipleParkingLotsOfSameAttendant(t *testing.T) {
	parkingLot1, parkingLot2, attendant1 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}
	attendant1.NewAttendant(&Strategy.BasicLotStrategy{})
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1)
	_, _ = attendant1.assign(parkingLot1), attendant1.assign(parkingLot2)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	firstCarTicket, _ := attendant1.park(firstCar)
	secondCarTicket, _ := attendant1.park(secondCar)

	_, _ = attendant1.UnPark(firstCarTicket)
	_, err := attendant1.UnPark(secondCarTicket)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUnparkMultipleCarsFromMultipleParkingLotsOfDifferentAttendant(t *testing.T) {
	parkingLot1, parkingLot2, attendant1, attendant2 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}, Attendant{}
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1)
	_, _ = attendant1.assign(parkingLot1), attendant2.assign(parkingLot2)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	firstCarTicket, _ := attendant1.park(firstCar)
	secondCarTicket, _ := attendant2.park(secondCar)

	_, _ = attendant1.UnPark(firstCarTicket)
	_, err := attendant2.UnPark(secondCarTicket)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUnparkCarsFromSameParkingLotThroughDifferentAttendants(t *testing.T) {
	parkingLot, attendant1, attendant2 := ParkingLot.ParkingLot{}, Attendant{}, Attendant{}
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _, _ = parkingLot.NewParkingLot(2), attendant1.assign(parkingLot), attendant2.assign(parkingLot)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	firstCarTicket, _ := attendant1.park(firstCar)
	secondCarTicket, _ := attendant1.park(secondCar)

	_, _ = attendant1.UnPark(firstCarTicket)
	_, err := attendant2.UnPark(secondCarTicket)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// ------------------------------- distributed parking Tests -------------------------------
func TestDistributedParking(t *testing.T) {
	setup()
	parkingLot2, parkingLot3, secondCar, thirdCar := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, &Car.Car{}, &Car.Car{}
	_, _ = parkingLot2.NewParkingLot(1), parkingLot3.NewParkingLot(1)
	secondCar, thirdCar = Car.NewCar("KA-01-HH-9999", Car.GREEN), Car.NewCar("KA-01-HH-1234", Car.BLUE)
	smartAttendant := Attendant{}
	smartAttendant.NewAttendant(&Strategy.SmartLotStrategy{})
	_, _, _ = smartAttendant.assign(parkingLot), smartAttendant.assign(parkingLot2), smartAttendant.assign(parkingLot3)

	_, _ = smartAttendant.park(car)       // parkingLot
	_, _ = smartAttendant.park(secondCar) // parkingLot2
	_, _ = smartAttendant.park(thirdCar)  // parkingLot3

	if !parkingLot.IsParkingLotFull() && parkingLot2.IsParkingLotFull() && parkingLot3.IsParkingLotFull() {
		t.Errorf("Expected first parkingLot have 1 slot free and other parkingLots to be full")
	}
}

func TestDistributedParkingWhenAllParkingLotsAreFull(t *testing.T) {
	parkingLot1, parkingLot2, parkingLot3 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}
	_, _, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1), parkingLot3.NewParkingLot(1)
	firstCar, secondCar, thirdCar, fourthCar := &Car.Car{}, &Car.Car{}, &Car.Car{}, &Car.Car{}
	firstCar, secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN), Car.NewCar("KA-01-HH-1234", Car.BLUE)
	thirdCar, fourthCar = Car.NewCar("KA-01-HH-7777", Car.RED), Car.NewCar("KA-01-HH-1111", Car.YELLOW)
	smartAttendant := Attendant{}
	smartAttendant.NewAttendant(&Strategy.SmartLotStrategy{})
	_, _, _ = smartAttendant.assign(parkingLot1), smartAttendant.assign(parkingLot2), smartAttendant.assign(parkingLot3)

	_, _ = smartAttendant.park(firstCar)  // parkingLot1
	_, _ = smartAttendant.park(secondCar) // parkingLot2
	_, _ = smartAttendant.park(thirdCar)  // parkingLot3
	_, err := smartAttendant.park(fourthCar)

	if err == nil {
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotFull, err)
	}
}

func TestDistributedParkingWhenAllParkingLotsAreFullAndUnparkCarAndParkAgainInSameParkingLot(t *testing.T) {
	parkingLot1, parkingLot2, parkingLot3 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}
	_, _, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1), parkingLot3.NewParkingLot(1)
	firstCar, secondCar, thirdCar, fourthCar := &Car.Car{}, &Car.Car{}, &Car.Car{}, &Car.Car{}
	firstCar, secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN), Car.NewCar("KA-01-HH-1234", Car.BLUE)
	thirdCar, fourthCar = Car.NewCar("KA-01-HH-7777", Car.RED), Car.NewCar("KA-01-HH-1111", Car.YELLOW)
	smartAttendant := Attendant{}
	smartAttendant.NewAttendant(&Strategy.SmartLotStrategy{})
	_, _, _ = smartAttendant.assign(parkingLot1), smartAttendant.assign(parkingLot2), smartAttendant.assign(parkingLot3)

	_, _ = smartAttendant.park(firstCar)        // parkingLot1
	ticket, _ := smartAttendant.park(secondCar) // parkingLot2
	_, _ = smartAttendant.park(thirdCar)        // parkingLot3
	_, err := smartAttendant.park(fourthCar)

	if err == nil {
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotFull, err)
	}

	_, _ = smartAttendant.UnPark(ticket)
	_, err = smartAttendant.park(fourthCar)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !parkingLot1.IsParkingLotFull() && !parkingLot2.IsParkingLotFull() && !parkingLot3.IsParkingLotFull() {
		t.Errorf("Expected all parkingLots to be full")
	}
}
