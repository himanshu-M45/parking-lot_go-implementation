package Car

type Car struct {
	registrationNumber string
	color              CarColor
	isParked           bool
}

func NewCar(registrationNumber string, color CarColor) Car {
	return Car{registrationNumber: registrationNumber, color: color, isParked: false}
}

func (car Car) IsIdenticalCar(registrationNumber string) bool {
	return car.registrationNumber == registrationNumber
}

func (car Car) IsSameColor(color CarColor) bool {
	return car.color == color
}
