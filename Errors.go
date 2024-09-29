package customError

import "errors"

var (
	ErrInvalidTicket                     = errors.New("invalid ticket")
	ErrCarAlreadyParked                  = errors.New("car is already parked")
	ErrCarNotParked                      = errors.New("car is not parked in slot")
	ErrParkingLotAlreadyAssigned         = errors.New("parking lot already assigned")
	ErrParkingLotFull                    = errors.New("parking lot is full")
	ErrSlotNumberShouldBeGreaterThanZero = errors.New("slot number should be greater than 0")
)
