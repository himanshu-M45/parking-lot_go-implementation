package main

type Car struct {
	registrationNumber string
	color              CarColor
	isParked           bool
}

func newCar(registrationNumber string, color CarColor) Car {
	return Car{registrationNumber: registrationNumber, color: color, isParked: false}
}

func (car Car) isIdenticalCar(registrationNumber string) bool {
	return car.registrationNumber == registrationNumber
}

func (car Car) isSameColor(color CarColor) bool {
	return car.color == color
}
