package ParkingLot

import (
	"errors"
	customError "parking-lot"
	"parking-lot/Car"
	"testing"
)

// ------------------------------- parkingLot tests -------------------------------
func TestCreateParkingLotWith0Slots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.NewParkingLot(0)

	if !errors.Is(err, customError.ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", customError.ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestCreateParkingLotWithNegativeSlots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.NewParkingLot(-12)

	if !errors.Is(err, customError.ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", customError.ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestParkingLotIsSame(t *testing.T) {
	parkingLot := ParkingLot{}
	_ = parkingLot.NewParkingLot(1)

	if !parkingLot.IsSameParkingLot(parkingLot) { // !True
		t.Errorf("Expected parking lots to be same")
	}
}

func TestNewParkingLotIsEmpty(t *testing.T) {
	parkingLot := ParkingLot{}
	_ = parkingLot.NewParkingLot(1)

	if parkingLot.IsParkingLotFull() { // True
		t.Errorf("Expected parking lot to be empty")
	}
}

func TestParkingLotIsNotSame(t *testing.T) {
	firstParkingLot, secondParkingLot := ParkingLot{}, ParkingLot{}

	_, _ = firstParkingLot.NewParkingLot(1), secondParkingLot.NewParkingLot(1)

	if firstParkingLot.IsSameParkingLot(secondParkingLot) { // False
		t.Errorf("Expected parking lots to be same")
	}
}

// ------------------------------- park car tests -------------------------------
func TestParkCar(t *testing.T) {
	parkingLot, car := ParkingLot{}, &Car.Car{}
	_, car = parkingLot.NewParkingLot(1), Car.NewCar("KA-01-HH-1234", Car.BLACK)

	if _, err := parkingLot.Park(car); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestParkingLotWithOneSlotIsFullWhenCarParked(t *testing.T) {
	parkingLot, firstCar, secondCar := ParkingLot{}, &Car.Car{}, &Car.Car{}
	_ = parkingLot.NewParkingLot(1)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.WHITE)
	secondCar = Car.NewCar("KA-01-HH-1235", Car.BLUE)

	_, _ = parkingLot.Park(firstCar)

	if _, err := parkingLot.Park(secondCar); !errors.Is(err, customError.ErrParkingLotFull) { // !True
		t.Errorf("Expected error '%v', got %v", customError.ErrParkingLotFull, err)
	}
}

func TestParkingLotWithTwoSlotsHaveOneCarParkedAndIsNotFull(t *testing.T) {
	parkingLot, firstCar := ParkingLot{}, &Car.Car{}
	_ = parkingLot.NewParkingLot(2)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.WHITE)

	_, _ = parkingLot.Park(firstCar)

	if parkingLot.IsParkingLotFull() {
		t.Errorf("Expected parking lot to not be full")
	}
}

func TestParkingLotWithOneSlotsHaveOneCarParkedAndIsFull(t *testing.T) {
	parkingLot, firstCar := ParkingLot{}, &Car.Car{}
	_ = parkingLot.NewParkingLot(1)
	firstCar = Car.NewCar("KA-01-HH-1234", Car.WHITE)

	_, _ = parkingLot.Park(firstCar)

	if !parkingLot.IsParkingLotFull() {
		t.Errorf("Expected parking lot to not be full")
	}
}
