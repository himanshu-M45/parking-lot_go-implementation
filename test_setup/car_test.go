package test_setup

import (
	"parking-lot/Car"
	"testing"
)

func TestCheckCreatedCarShouldMatchItself(t *testing.T) {
	car := Car.Construct("KA-01-HH-1234", Car.RED)

	if !car.IsIdenticalCar("KA-01-HH-1234") {
		t.Errorf("Expected car to be created with registration number KA-01-HH-1234")
	}
}

func TestCheckCarColorIsSame(t *testing.T) {
	car := Car.Construct("RJ-14-HH-6294", Car.BLACK)

	if car.IsSameColor(Car.BLACK) == false {
		t.Errorf("Expected car to be RED")
	}
}

func TestCheckCarColorIsNotSame(t *testing.T) {
	car := Car.Construct("RJ-14-HH-6294", Car.BLUE)

	if car.IsSameColor(Car.BLACK) == true {
		t.Errorf("Expected car to be RED")
	}
}
