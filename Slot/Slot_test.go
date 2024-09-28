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
