package tests

import (
	"github.com/stretchr/testify/assert"
	"parking-lot/Car"
	"parking-lot/common/custom_errors"
	"testing"
)

func TestSlotInitialization(t *testing.T) {
	Setup()
	assert.False(t, slot.IsSlotOccupied())
}

// ------------------------------- park tests -------------------------------
func TestParkCarInSlot(t *testing.T) {
	Setup()

	carTicket, _ := slot.Park(car)

	assert.NotNil(t, carTicket)
}

func TestParkCarAlreadyParkedCar(t *testing.T) {
	Setup()

	_, _ = slot.Park(car)
	_, err := slot.Park(car)

	assert.Equal(t, custom_errors.ErrCarAlreadyParked, err)
}

func TestCannotParkSameCarTwice(t *testing.T) {
	Setup()

	_, _ = slot.Park(car)
	_, err := slot.Park(car)

	assert.Equal(t, custom_errors.ErrCarAlreadyParked, err)
	assert.True(t, car.IsCarParked())
}

// ------------------------------- check parked Car color tests -------------------------------
func TestCheckBlackColorCarIsParkedInSlot(t *testing.T) {
	Setup()
	_, _ = slot.Park(car)
	assert.True(t, slot.IsCarColor(Car.BLACK))
}

func TestCheckBlackColorCarIsNotParkedInSlot(t *testing.T) {
	Setup()
	_, _ = slot.Park(car3)
	assert.False(t, slot.IsCarColor(Car.BLACK))
}

// ------------------------------- check Car by reg num tests -------------------------------
func TestGetTicketIfCarMatches(t *testing.T) {
	Setup()

	_, _ = slot.Park(car)
	ticket, _ := slot.GetTicketIfCarMatches("KA-01-HH-1234")

	assert.True(t, ticket.ValidateTicket(ticket))
}

func TestCarNotFoundErrIfCarIsNotAvailable(t *testing.T) {
	Setup()

	_, err := slot.GetTicketIfCarMatches("KA-01-HH-1235")

	assert.Equal(t, custom_errors.ErrCarNotParked, err)
}

// ------------------------------- unpark tests -------------------------------
func TestUnparkCarFromSlot(t *testing.T) {
	Setup()

	ticket, _ := slot.Park(car)
	receivedCar, _ := slot.UnPark(ticket)

	if !receivedCar.IsIdenticalCar("KA-01-HH-1234") {
		t.Errorf("Expected Car to be unparked")
	}
	assert.True(t, receivedCar.IsIdenticalCar("KA-01-HH-1234"))
}

func TestUnparkAlreadyUnParkedCar(t *testing.T) {
	Setup()

	ticket, _ := slot.Park(car)
	_, _ = slot.UnPark(ticket)

	_, err := slot.UnPark(ticket)

	assert.Equal(t, custom_errors.ErrInvalidTicket, err)
}
