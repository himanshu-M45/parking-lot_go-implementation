package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Custom errors
var (
	ErrParkingLotFull                    = errors.New("parking lot is full")
	ErrSlotNumberShouldBeGreaterThanZero = errors.New("slot number should be greater than 0")
)

type ParkingLot struct {
	isFull       bool
	slots        []Slot
	parkingLotId string
}

func (parkingLot *ParkingLot) newParkingLot(numberOfSlots int) error {
	if numberOfSlots < 1 {
		return ErrSlotNumberShouldBeGreaterThanZero
	}
	parkingLot.isFull = false
	parkingLot.slots = make([]Slot, numberOfSlots)
	parkingLot.parkingLotId = fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
	fmt.Println(parkingLot.parkingLotId)
	return nil
}

func (parkingLot *ParkingLot) isSameParkingLot(receivedParkingLot ParkingLot) bool {
	return parkingLot.parkingLotId == receivedParkingLot.parkingLotId
}
