package main

import (
	"errors"
	"testing"
)

// ------------------------------- assign parkingLot tests -------------------------------
func TestAssignParkingLotToAttendant(t *testing.T) {
	parkingLot := ParkingLot{}
	if err := parkingLot.newParkingLot(1); err != nil { // nil
		t.Fatalf("Expected no error, got %v", err)
	}

	attendant := Attendant{}
	err := attendant.assign(parkingLot)
	if err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAssignTwoParkingLotToAttendant(t *testing.T) {
	parkingLot1 := ParkingLot{}
	if err := parkingLot1.newParkingLot(1); err != nil { // nil
		t.Fatalf("Expected no error, got %v", err)
	}
	parkingLot2 := ParkingLot{}
	if err := parkingLot2.newParkingLot(1); err != nil { // nil
		t.Fatalf("Expected no error, got %v", err)
	}
	attendant := Attendant{}
	if err := attendant.assign(parkingLot1); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
	if err := attendant.assign(parkingLot2); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCannotAssignSameParkingLotTwice(t *testing.T) {
	parkingLot := ParkingLot{}
	if err := parkingLot.newParkingLot(1); err != nil { // nil
		t.Fatalf("Expected no error, got %v", err)
	}

	attendant := Attendant{}
	if err := attendant.assign(parkingLot); err != nil { // nil
		t.Errorf("Expected no error, got %v", err)
	}

	err := attendant.assign(parkingLot)
	if !errors.Is(err, ErrParkingLotAlreadyAssigned) { // !True
		t.Errorf("Expected error '%v', got %v", ErrParkingLotAlreadyAssigned, err)
	}
}

func TestAssignSameParkingLotToMultipleAttendants(t *testing.T) {
	parkingLot := ParkingLot{}
	if err := parkingLot.newParkingLot(1); err != nil { // nil
		t.Fatalf("Expected no error, got %v", err)
	}

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
