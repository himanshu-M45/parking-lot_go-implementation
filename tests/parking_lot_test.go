package tests

import (
	"github.com/stretchr/testify/assert"
	"parking-lot/Car"
	"parking-lot/common/custom_errors"
	"parking-lot/parking_lot"
	"parking-lot/receipt"
	"testing"
)

// ------------------------------- parkingLot tests -------------------------------
func TestCreateParkingLotWith0Slots(t *testing.T) {
	Setup()
	_, err := owner.CreateParkingLot(0)
	assert.Equal(t, custom_errors.ErrSlotNumberShouldBeGreaterThanZero, err)
}

func TestCreateParkingLotWithoutOwner(t *testing.T) {
	anotherParkingLot := parking_lot.ParkingLot{}
	err := anotherParkingLot.Construct(1, "")
	assert.Equal(t, custom_errors.ErrCannotCreateParkingLotWithoutOwner, err)
}

func TestCreateParkingLotWithNegativeSlots(t *testing.T) {
	Setup()
	_, err := owner.CreateParkingLot(-3)
	assert.Equal(t, custom_errors.ErrSlotNumberShouldBeGreaterThanZero, err)
}

func TestNewParkingLotIsEmpty(t *testing.T) {
	Setup()
	assert.False(t, parkingLot.IsParkingLotFull())
}

func TestParkingLotIsSame(t *testing.T) {
	Setup()
	assert.True(t, parkingLot.IsSameParkingLot(parkingLot))
}

func TestParkingLotIsNotSame(t *testing.T) {
	Setup()
	assert.False(t, parkingLot.IsSameParkingLot(parkingLot2))
}

// ------------------------------- park tests -------------------------------
func TestParkCar(t *testing.T) {
	Setup()
	_, err := parkingLot.Park(car)

	assert.NoError(t, err)
}

func TestParkSecondCarInParkingLotWithOneSlot(t *testing.T) {
	Setup()
	_, _ = parkingLot.Park(car)
	_, err := parkingLot.Park(car1)

	assert.Equal(t, custom_errors.ErrParkingLotFull, err)
}

func TestParkingLotWithTwoSlotsHaveOneCarParkedAndIsNotFull(t *testing.T) {
	Setup()
	_, _ = parkingLot2.Park(car)

	if parkingLot2.IsParkingLotFull() {
		t.Errorf("Expected parking lot to not be full")
	}
	assert.False(t, parkingLot2.IsParkingLotFull())
}

func TestParkingLotWithOneSlotsHaveOneCarParkedAndIsFull(t *testing.T) {
	Setup()
	_, _ = parkingLot.Park(car)

	assert.True(t, parkingLot.IsParkingLotFull())
}

// ------------------------------- count Car by color tests -------------------------------
func TestGetCountOfWhiteColorCars(t *testing.T) {
	Setup()
	_, _ = parkingLot4.Park(car)
	_, _ = parkingLot4.Park(car1)
	_, _ = parkingLot4.Park(car2)

	assert.Equal(t, 1, parkingLot4.CountCarsByColor(Car.WHITE))
}

func TestGetCountOfBlackColorCars(t *testing.T) {
	Setup()
	_, _ = parkingLot4.Park(car)
	_, _ = parkingLot4.Park(car1)
	_, _ = parkingLot4.Park(car2)

	assert.Equal(t, 2, parkingLot4.CountCarsByColor(Car.BLACK))
}

// ------------------------------- check Car by reg num tests -------------------------------
func TestCheckTheGivenCarIsAvailableInParkingLot(t *testing.T) {
	Setup()
	_, _ = parkingLot.Park(car)
	receivedTicket, _ := parkingLot.GetCarParkedInfoByRegNo("KA-01-HH-1234")

	assert.True(t, receivedTicket.ValidateTicket(receivedTicket))
}

func TestCheckTheGivenCarIsNotAvailableInParkingLot(t *testing.T) {
	Setup()
	_, err := parkingLot.GetCarParkedInfoByRegNo("KA-01-HH-1235")

	assert.Equal(t, custom_errors.ErrCarNotParked, err)
}

// ------------------------------- unpark tests -------------------------------
func TestUnparkCarFromParkingLot(t *testing.T) {
	Setup()
	receivedTicket, _ := parkingLot.Park(car)
	_, err := parkingLot.UnPark(receivedTicket)

	assert.NoError(t, err)
}

func TestCannotUnparkCarFromParkingLotWithInvalidTicket(t *testing.T) {
	Setup()
	receivedTicket, _ := parkingLot.Park(car)
	_, _ = parkingLot.UnPark(receivedTicket)
	_, err := parkingLot.UnPark(receivedTicket)

	assert.Equal(t, custom_errors.ErrInvalidTicket, err)
}

func TestCannotUnparkUnavailableCarFromParkingLot(t *testing.T) {
	Setup()
	dummyTicket := receipt.Receipt{}
	dummyTicket = *receipt.Construct()

	_, err := parkingLot.UnPark(dummyTicket)

	assert.Equal(t, custom_errors.ErrInvalidTicket, err)
}
