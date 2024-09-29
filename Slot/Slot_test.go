package Slot

import (
	"errors"
	customError "parking-lot"
	"parking-lot/Car"
	"parking-lot/Ticket"
	"testing"
)

func TestSlotInitialization(t *testing.T) {
	slot := Slot{}
	slot.NewSlot()
	if slot.IsSlotOccupied() {
		t.Errorf("Slot should be empty")
	}
}

// ------------------------------- park Tests -------------------------------
func TestParkCar(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.WHITE)

	ticket, _ := slot.Park(car)
	if ticket == (Ticket.Ticket{}) {
		t.Errorf("Ticket should be returned")
	}
}

func TestParkCarInOccupiedSlot(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.WHITE)

	_, _ = slot.Park(car)
	if _, err := slot.Park(car); !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
}

func TestParkSameCarTwice(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.WHITE)

	_, _ = slot.Park(car)
	if _, err := slot.Park(car); !errors.Is(err, customError.ErrCarAlreadyParked) {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarAlreadyParked, err)
	}
	if !car.IsCarParked() {
		t.Errorf("Expected car to be parked")
	}
}

// ------------------------------- unpark Tests -------------------------------

// ------------------------------- check parked car color Tests -------------------------------
func TestCheckBlackColorCarIsParkedInSlot(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.BLACK)

	_, _ = slot.Park(car)
	if !slot.IsCarColor(Car.BLACK) {
		t.Errorf("Expected car color to be black")
	}
}

func TestCheckBlackColorCarIsNotParkedInSlot(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.RED)

	_, _ = slot.Park(car)
	if slot.IsCarColor(Car.BLACK) {
		t.Errorf("Expected car color to be red")
	}
}

// ------------------------------- check car by reg num Tests -------------------------------
func TestGetTicketIfCarMatches(t *testing.T) {
	slot, car := Slot{}, &Car.Car{}
	slot.NewSlot()
	car = Car.NewCar("KA-01-HH-1234", Car.BLACK)

	_, _ = slot.Park(car)
	ticket, _ := slot.GetTicketIfCarMatches("KA-01-HH-1234")

	if !ticket.ValidateTicket(ticket) {
		t.Errorf("Expected ticket to match")
	}
}

func TestCarNotFoundErrIfCarIsNotAvailable(t *testing.T) {
	slot := Slot{}
	slot.NewSlot()

	_, err := slot.GetTicketIfCarMatches("KA-01-HH-1235")

	if err == nil {
		t.Errorf("Expected error '%v', got %v", customError.ErrCarNotParked, err)
	}
}
