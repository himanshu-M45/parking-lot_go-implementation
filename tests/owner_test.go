package tests

import (
	"github.com/stretchr/testify/assert"
	"parking-lot/common/custom_errors"
	"testing"
)

func TestOwner(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		assert.NotNil(t, &owner)
	})

	t.Run("CreateParkingLot", func(t *testing.T) {
		t.Run("SingleParkingLot", func(t *testing.T) {
			assert.NotNil(t, parkingLot)
		})

		t.Run("MultipleParkingLots", func(t *testing.T) {
			assert.NotNil(t, parkingLot)
			assert.NotNil(t, parkingLot1)
			assert.NotNil(t, parkingLot2)
		})
	})

	t.Run("AttendantTest", func(t *testing.T) {
		t.Run("AssignOwnedParkingLotToAttendant", func(t *testing.T) {
			Setup()
			assert.NoError(t, owner.Assign(parkingLot, &attendant))
		})

		t.Run("AssignToNotOwnedParkingLot", func(t *testing.T) {
			err := owner.Assign(parkingLot3, &attendant)
			assert.Equal(t, custom_errors.ErrOwnerDoesNotOwnParkingLot, err)
		})

		t.Run("AssignMultipleAttendantsToSameParkingLot", func(t *testing.T) {
			Setup()
			assert.NoError(t, owner.Assign(parkingLot, &attendant))
			assert.NoError(t, owner.Assign(parkingLot, &newAttendant))
		})

		t.Run("AssignSameAttendantToMultipleParkingLots", func(t *testing.T) {
			Setup()
			assert.NoError(t, owner.Assign(parkingLot, &attendant))
			assert.NoError(t, owner.Assign(parkingLot1, &attendant))
		})

		t.Run("AssignAlreadyAssignedSmartAttendant", func(t *testing.T) {
			_ = owner.Assign(parkingLot, &smartAttendant)
			err := owner.Assign(parkingLot, &smartAttendant)
			assert.Equal(t, custom_errors.ErrParkingLotAlreadyAssigned, err)
		})

		t.Run("AssignOwnedParkingLotToOwner", func(t *testing.T) {
			Setup()
			err := owner.Assign(parkingLot, owner.Attendant)
			assert.NoError(t, err)
		})
	})

	t.Run("ParkTest", func(t *testing.T) {
		t.Run("ParkCarThroughOwner", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, owner.Attendant)
			ticket, err := owner.Park(car)
			assert.NoError(t, err)
			assert.True(t, ticket.ValidateTicket(ticket))
		})

		t.Run("ParkCarThroughOwnerWithMultipleParkingLots", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, owner.Attendant)
			_ = owner.Assign(parkingLot1, owner.Attendant)

			firstTicket, _ := owner.Park(car)
			secondTicket, _ := owner.Park(car1)

			assert.True(t, firstTicket.ValidateTicket(firstTicket))
			assert.True(t, secondTicket.ValidateTicket(secondTicket))
		})

		t.Run("ParkAlreadyParkedCarThroughOwner", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot2, owner.Attendant)

			_, _ = owner.Park(car)
			_, err := owner.Park(car)

			assert.Equal(t, custom_errors.ErrCarAlreadyParked, err)
		})

		t.Run("BasicAttendantParkCarThroughOwner", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, &attendant)

			_, err := attendant.Park(car)

			assert.NoError(t, err)
		})

		t.Run("TestParkCarBySmartAttendantThroughOwner", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, &smartAttendant)

			_, err := smartAttendant.Park(car)

			assert.NoError(t, err)
		})

	})

	t.Run("UnParkTest", func(t *testing.T) {
		t.Run("UnparkCarThroughOwner", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, owner.Attendant)

			ticket, _ := owner.Park(car)
			_, err := owner.UnPark(ticket)

			assert.NoError(t, err)
		})

		t.Run("UnparkCarThroughOwnerWithMultipleParkingLots", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, owner.Attendant)
			_ = owner.Assign(parkingLot1, owner.Attendant)

			firstTicket, _ := owner.Park(car)
			secondTicket, _ := owner.Park(car1)

			_, err1 := owner.UnPark(firstTicket)
			assert.NoError(t, err1)
			_, err2 := owner.UnPark(secondTicket)
			assert.NoError(t, err2)
		})

		t.Run("UnparkAlreadyUnParkedCar", func(t *testing.T) {
			Setup()
			_ = owner.Assign(parkingLot, owner.Attendant)

			ticket, _ := owner.Park(car)
			_, _ = owner.UnPark(ticket)
			_, err := owner.UnPark(ticket)

			assert.Equal(t, custom_errors.ErrInvalidTicket, err)
		})
	})

	t.Run("NotifiableTest", func(t *testing.T) {})
}
