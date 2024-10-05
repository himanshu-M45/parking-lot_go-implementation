package tests

import (
	"github.com/stretchr/testify/assert"
	"parking-lot/common/custom_errors"
	"testing"
)

func TestOwnerInitialization(t *testing.T) {
	assert.NotNil(t, &owner)
}

func TestCreateParkingLotToOwner(t *testing.T) {
	Setup()
	assert.NotNil(t, parkingLot)
}

func TestCreateMultipleParkingLotsToOwner(t *testing.T) {
	Setup()
	assert.NotNil(t, parkingLot)
	assert.NotNil(t, parkingLot1)
	assert.NotNil(t, parkingLot2)
}

// ---------------------------- role tests ----------------------------
func TestAssignOwnedParkingLotToAttendant(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
}

func TestAssignToNotOwnedParkingLot(t *testing.T) {
	Setup()
	err := owner.Assign(parkingLot3, &attendant)
	assert.Equal(t, custom_errors.ErrOwnerDoesNotOwnParkingLot, err)
}

func TestAssignMultipleAttendantsToSameParkingLot(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
	assert.NoError(t, owner.Assign(parkingLot, &newAttendant))
}

func TestAssignSameAttendantToMultipleParkingLots(t *testing.T) {
	Setup()
	assert.NoError(t, owner.Assign(parkingLot, &attendant))
	assert.NoError(t, owner.Assign(parkingLot1, &attendant))
}

func TestAssignAlreadyAssignedSmartAttendant(t *testing.T) {
	Setup()
	_ = owner.Assign(parkingLot, &smartAttendant)
	err := owner.Assign(parkingLot, &smartAttendant)
	assert.Equal(t, custom_errors.ErrParkingLotAlreadyAssigned, err)
}
