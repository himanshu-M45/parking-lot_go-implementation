package customError

import "errors"

var (
	ErrParkingLotFull                    = errors.New("parking lot is full")
	ErrSlotNumberShouldBeGreaterThanZero = errors.New("slot number should be greater than 0")
	ErrCarAlreadyParked                  = errors.New("car is already parked")
	ErrParkingLotAlreadyAssigned         = errors.New("parking lot already assigned")
)
