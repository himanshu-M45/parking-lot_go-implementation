package ParkingLot

import (
	"errors"
	"testing"
)

func TestCreateParkingLotWith0Slots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.NewParkingLot(0)

	if !errors.Is(err, ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestCreateParkingLotWithNegativeSlots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.NewParkingLot(-12)

	if !errors.Is(err, ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestParkingLotIsSame(t *testing.T) {
	parkingLot := ParkingLot{}
	if err := parkingLot.NewParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	if !parkingLot.IsSameParkingLot(parkingLot) { // !True
		t.Errorf("Expected parking lots to be same")
	}
}

func TestParkingLotIsNotSame(t *testing.T) {
	firstParkingLot := ParkingLot{}
	if err := firstParkingLot.NewParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
	secondParkingLot := ParkingLot{}
	if err := secondParkingLot.NewParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	if firstParkingLot.IsSameParkingLot(secondParkingLot) { // False
		t.Errorf("Expected parking lots to be same")
	}
}
