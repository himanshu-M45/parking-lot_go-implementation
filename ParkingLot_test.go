package main

import (
	"errors"
	"testing"
)

func TestCreateParkingLotWith0Slots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.newParkingLot(0)

	if !errors.Is(err, ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestCreateParkingLotWithNegativeSlots(t *testing.T) {
	parkingLot := ParkingLot{}

	err := parkingLot.newParkingLot(-12)

	if !errors.Is(err, ErrSlotNumberShouldBeGreaterThanZero) { // !True
		t.Errorf("Expected error '%v', got %v", ErrSlotNumberShouldBeGreaterThanZero, err)
	}
}

func TestParkingLotIsSame(t *testing.T) {
	parkingLot := ParkingLot{}
	if err := parkingLot.newParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	if !parkingLot.isSameParkingLot(parkingLot) { // !True
		t.Errorf("Expected parking lots to be same")
	}
}

func TestParkingLotIsNotSame(t *testing.T) {
	firstParkingLot := ParkingLot{}
	if err := firstParkingLot.newParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
	secondParkingLot := ParkingLot{}
	if err := secondParkingLot.newParkingLot(1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	if firstParkingLot.isSameParkingLot(secondParkingLot) { // False
		t.Errorf("Expected parking lots to be same")
	}
}
