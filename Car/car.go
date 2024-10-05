package Car

type Car struct {
	registeredNumber string
	color            CarColor
	isParked         bool
}

func Construct(registeredNumber string, color CarColor) *Car {
	return &Car{registeredNumber: registeredNumber, color: color, isParked: false}
}

func (car *Car) IsIdenticalCar(registeredNumber string) bool {
	return car.registeredNumber == registeredNumber
}

func (car *Car) IsSameColor(color CarColor) bool {
	return car.color == color
}

func (car *Car) IsCarParked() bool { return car.isParked }

func (car *Car) SetCarParked(isParked bool) { car.isParked = isParked }
