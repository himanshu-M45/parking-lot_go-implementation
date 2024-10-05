package customError

import "errors"

var (
	ErrCarAlreadyParked                  = errors.New("Car is already parked")
	ErrCarNotParked                      = errors.New("Car is not parked in slots")
	ErrInvalidTicket                     = errors.New("invalid receipt")
	ErrOwnerDoesNotOwnParkingLot         = errors.New("owner does not own current parking lot")
	ErrParkingLotAlreadyAssigned         = errors.New("parking lot already assigned")
	ErrParkingLotFull                    = errors.New("parking lot is full")
	ErrSlotNumberShouldBeGreaterThanZero = errors.New("slots number should be greater than 0")
)
