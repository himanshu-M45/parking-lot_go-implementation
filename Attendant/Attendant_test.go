package Attendant

import (
	"errors"
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/ParkingLot"
	"testing"
)

// ------------------------------- assign parkingLot tests -------------------------------
func TestAssignParkingLotToAttendant(t *testing.T) {
	parkingLot := ParkingLot.ParkingLot{}
	_ = parkingLot.NewParkingLot(1)

	attendant := Attendant{}
	err := attendant.assign(parkingLot)
	if err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAssignTwoParkingLotToAttendant(t *testing.T) {
	parkingLot1 := ParkingLot.ParkingLot{}
	_ = parkingLot1.NewParkingLot(1)
	parkingLot2 := ParkingLot.ParkingLot{}
	_ = parkingLot2.NewParkingLot(1)

	attendant := Attendant{}
	if err := attendant.assign(parkingLot1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
	if err := attendant.assign(parkingLot2); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCannotAssignSameParkingLotTwice(t *testing.T) {
	parkingLot := ParkingLot.ParkingLot{}
	_ = parkingLot.NewParkingLot(1)
	attendant := Attendant{}
	_ = attendant.assign(parkingLot)

	err := attendant.assign(parkingLot)
	if !errors.Is(err, customError.ErrParkingLotAlreadyAssigned) { // !True
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotAlreadyAssigned, err)
	}
}

func TestAssignSameParkingLotToMultipleAttendants(t *testing.T) {
	parkingLot := ParkingLot.ParkingLot{}
	_ = parkingLot.NewParkingLot(1)

	attendant1 := Attendant{}
	if err := attendant1.assign(parkingLot); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	attendant2 := Attendant{}
	if err := attendant2.assign(parkingLot); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

// ------------------------------- park through attendant tests -------------------------------
func TestParkingCarThroughAttendant(t *testing.T) {
	parkingLot, attendant, car := ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}
	_, _ = parkingLot.NewParkingLot(1), attendant.assign(parkingLot)
	car = Car.NewCar("KA-01-HH-1234", Car.YELLOW)

	_, err := attendant.park(car)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestTryToParkCarWhenParkingLotIsFull(t *testing.T) {
	parkingLot, attendant, firstCar, secondCar := ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}, &Car.Car{}
	_, _ = parkingLot.NewParkingLot(1), attendant.assign(parkingLot)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant.park(firstCar)
	_, err := attendant.park(secondCar)

	if !errors.Is(err, customError.ErrParkingLotFull) {
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotFull, err)
	}
}

func TestParkSameCarTwice(t *testing.T) {
	parkingLot, attendant, car := ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}
	_, _ = parkingLot.NewParkingLot(2), attendant.assign(parkingLot)
	car = Car.NewCar("KA-01-HH-1234", Car.YELLOW)

	_, _ = attendant.park(car)
	_, err := attendant.park(car)

	if !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
}

func TestParkMultipleCarsInParkingLotThroughAttendant(t *testing.T) {
	parkingLot, attendant, firstCar, secondCar := ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}, &Car.Car{}
	_, _ = parkingLot.NewParkingLot(2), attendant.assign(parkingLot)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant.park(firstCar)
	_, err := attendant.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkCarInMultipleParkingLotsThroughSameAttendant(t *testing.T) {
	parkingLot1, parkingLot2, attendant := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(2), parkingLot2.NewParkingLot(1)
	_, _ = attendant.assign(parkingLot1), attendant.assign(parkingLot2)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant.park(firstCar)
	_, err := attendant.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkSameCarInDifferentParkingLotsOfSameAttendant(t *testing.T) {
	parkingLot1, parkingLot2, attendant, car := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}, &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(2), parkingLot2.NewParkingLot(1)
	_, _ = attendant.assign(parkingLot1), attendant.assign(parkingLot2)
	car = Car.NewCar("KA-01-HH-1234", Car.YELLOW)

	_, _ = attendant.park(car)
	_, err := attendant.park(car)

	if !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
}

func TestParkCarInDifferentParkingLotThroughDifferentAttendants(t *testing.T) {
	parkingLot1, parkingLot2, attendant1, attendant2 := ParkingLot.ParkingLot{}, ParkingLot.ParkingLot{}, Attendant{}, Attendant{}
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _ = parkingLot1.NewParkingLot(1), parkingLot2.NewParkingLot(1)
	_, _ = attendant1.assign(parkingLot1), attendant2.assign(parkingLot2)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.YELLOW)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.GREEN)

	_, _ = attendant1.park(firstCar)
	_, err := attendant2.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkCarInSameParkingLotThroughDifferentAttendants(t *testing.T) {
	parkingLot, attendant1, attendant2 := ParkingLot.ParkingLot{}, Attendant{}, Attendant{}
	firstCar, secondCar := &Car.Car{}, &Car.Car{}
	_, _, _ = parkingLot.NewParkingLot(2), attendant1.assign(parkingLot), attendant2.assign(parkingLot)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.BLUE)
	secondCar = Car.NewCar("KA-01-HH-9999", Car.BLUE)

	_, _ = attendant1.park(firstCar)
	_, err := attendant2.park(secondCar)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// ------------------------------- unpark through attendant Tests -------------------------------
