package customError

import "errors"

var (
	ErrCarAlreadyParked                  = errors.New("car is already parked")
	ErrCarNotParked                      = errors.New("car is not parked in slot")
	ErrInvalidTicket                     = errors.New("invalid ticket")
	ErrOwnerDoesNotOwnParkingLot         = errors.New("owner does not own current parking lot")
	ErrParkingLotAlreadyAssigned         = errors.New("parking lot already assigned")
	ErrParkingLotFull                    = errors.New("parking lot is full")
	ErrSlotNumberShouldBeGreaterThanZero = errors.New("slot number should be greater than 0")
)
