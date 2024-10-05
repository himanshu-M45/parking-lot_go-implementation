package tests

import (
	"github.com/stretchr/testify/assert"
	"parking-lot/common/custom_errors"
	"testing"
)

// ------------------------------- assign parkingLot tests -------------------------------
func TestAssignParkingLotToAttendant(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
}

func TestAssignTwoParkingLotToAttendant(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
	assert.NoError(t, owner.Assign(parkingLot1, &attendant))
}

func TestCannotAssignSameParkingLotTwice(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	err := owner.Assign(parkingLot, &attendant)
	assert.Equal(t, custom_errors.ErrParkingLotAlreadyAssigned, err)
}

func TestAssignSameParkingLotToMultipleAttendants(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
	assert.NoError(t, owner.Assign(parkingLot, &newAttendant))
}

func TestAssignMultipleParkingLotsToSameAttendant(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
	assert.NoError(t, owner.Assign(parkingLot1, &attendant))
}

// ------------------------------- park through attendant tests -------------------------------
func TestParkingCarThroughAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)

	_, err := attendant.Park(car)

	assert.NoError(t, err)
}

func TestTryToParkCarWhenParkingLotIsFull(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot2, &attendant)

	_, _ = attendant.Park(car)
	_, _ = attendant.Park(car1)
	_, err := attendant.Park(car2)

	assert.Equal(t, custom_errors.ErrParkingLotFull, err)
}

func TestParkSameCarTwiceExpectError(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot2, &attendant)

	_, _ = attendant.Park(car)
	_, err := attendant.Park(car)

	assert.Equal(t, custom_errors.ErrCarAlreadyParked, err)
}

func TestParkMultipleCarsInParkingLotThroughAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot2, &attendant)

	_, _ = attendant.Park(car)
	_, err := attendant.Park(car1)

	assert.NoError(t, err)
}

func TestParkCarInMultipleParkingLotsThroughSameAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &attendant)

	_, _ = attendant.Park(car)
	_, err := attendant.Park(car1)

	assert.NoError(t, err)
}

func TestParkSameCarInDifferentParkingLotsOfSameAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &attendant)

	_, _ = attendant.Park(car)
	_, err := attendant.Park(car)

	assert.Equal(t, custom_errors.ErrCarAlreadyParked, err)
}

func TestParkCarInDifferentParkingLotThroughDifferentAttendants(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &newAttendant)

	_, _ = attendant.Park(car)
	_, err := newAttendant.Park(car1)

	assert.NoError(t, err)
}

func TestParkCarInSameParkingLotThroughDifferentAttendants(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot2, &attendant)
	_ = owner.Assign(parkingLot2, &newAttendant)

	_, _ = attendant.Park(car)
	_, err := newAttendant.Park(car1)

	assert.NoError(t, err)
}

// ------------------------------- unpark through attendant tests -------------------------------
func TestUnParkCarThroughAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)

	ticket, _ := attendant.Park(car)
	receivedCar, _ := attendant.UnPark(ticket)

	assert.True(t, receivedCar.IsIdenticalCar("KA-01-HH-1234"))
}

func TestUnParkCarFromMultipleParkingLots(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &attendant)

	firstCarTicket, _ := attendant.Park(car)
	secondCarTicket, _ := attendant.Park(car1)
	_, _ = attendant.UnPark(firstCarTicket)
	_, err := attendant.UnPark(secondCarTicket)

	assert.NoError(t, err)
}

func TestUnParkAlreadyUnParkedCar(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)

	ticket, _ := attendant.Park(car)
	_, _ = attendant.UnPark(ticket)
	_, err := attendant.UnPark(ticket)

	assert.Equal(t, custom_errors.ErrInvalidTicket, err)
}

func TestUnParkMultipleCarsFromMultipleParkingLotsOfSameAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &attendant)

	firstCarTicket, _ := attendant.Park(car)
	secondCarTicket, _ := attendant.Park(car1)

	_, _ = attendant.UnPark(firstCarTicket)
	_, err := attendant.UnPark(secondCarTicket)

	assert.NoError(t, err)
}

func TestUnparkMultipleCarsFromMultipleParkingLotsOfDifferentAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &attendant)
	_ = owner.Assign(parkingLot1, &newAttendant)

	firstCarTicket, _ := attendant.Park(car)
	secondCarTicket, _ := newAttendant.Park(car1)

	_, _ = attendant.UnPark(firstCarTicket)
	_, err := newAttendant.UnPark(secondCarTicket)

	assert.NoError(t, err)
}

func TestUnparkCarsFromSameParkingLotThroughDifferentAttendants(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot2, &attendant)
	_ = owner.Assign(parkingLot2, &newAttendant)

	firstCarTicket, _ := attendant.Park(car)
	secondCarTicket, _ := newAttendant.Park(car1)

	_, _ = attendant.UnPark(firstCarTicket)
	_, err := newAttendant.UnPark(secondCarTicket)

	assert.NoError(t, err)
}

// ------------------------------- distributed parking tests -------------------------------
func TestDistributedParking(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &smartAttendant)
	_ = owner.Assign(parkingLot1, &smartAttendant)
	_ = owner.Assign(parkingLot2, &smartAttendant)

	_, _ = smartAttendant.Park(car)  // parkingLot
	_, _ = smartAttendant.Park(car1) // parkingLot2
	_, _ = smartAttendant.Park(car2) // parkingLot3

	assert.True(t, parkingLot.IsParkingLotFull())
	assert.True(t, parkingLot1.IsParkingLotFull())
	assert.True(t, !parkingLot2.IsParkingLotFull())
}

func TestDistributedParkingWhenAllParkingLotsAreFull(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &smartAttendant)
	_ = owner.Assign(parkingLot1, &smartAttendant)
	_ = owner.Assign(parkingLot3, &smartAttendant)

	_, _ = smartAttendant.Park(car)  // parkingLot1
	_, _ = smartAttendant.Park(car1) // parkingLot2
	_, _ = smartAttendant.Park(car2) // parkingLot3
	_, err := smartAttendant.Park(car3)

	assert.Equal(t, custom_errors.ErrParkingLotFull, err)
}

func TestDistributedParkingWhenAllParkingLotsAreFullAndUnparkCarAndParkAgainInSameParkingLot(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &smartAttendant)
	_ = owner.Assign(parkingLot1, &smartAttendant)
	_ = owner.Assign(parkingLot3, &smartAttendant)

	_, _ = smartAttendant.Park(car)        // parkingLot1
	ticket, _ := smartAttendant.Park(car1) // parkingLot2
	_, _ = smartAttendant.Park(car2)       // parkingLot3
	_, err := smartAttendant.Park(car3)

	assert.Equal(t, custom_errors.ErrParkingLotFull, err)

	_, _ = smartAttendant.UnPark(ticket)
	_, err = smartAttendant.Park(car3)

	assert.NoError(t, err)

	assert.True(t, parkingLot.IsParkingLotFull())
	assert.True(t, parkingLot1.IsParkingLotFull())
	assert.False(t, parkingLot3.IsParkingLotFull())
}
