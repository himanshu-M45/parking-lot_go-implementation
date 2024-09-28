package Car

import (
	"testing"
)

func TestCheckCreatedCarShouldMatchItself(t *testing.T) {
	car := NewCar("KA-01-HH-1234", RED)

	if !car.IsIdenticalCar("KA-01-HH-1234") {
		t.Errorf("Expected car to be created with registration number KA-01-HH-1234")
	}
}

func TestCheckCarColorIsSame(t *testing.T) {
	car := NewCar("RJ-14-HH-6294", BLACK)

	if car.IsSameColor(BLACK) == false {
		t.Errorf("Expected car to be RED")
	}
}

func TestCheckCarColorIsNotSame(t *testing.T) {
	car := NewCar("RJ-14-HH-6294", BLUE)

	if car.IsSameColor(BLACK) == true {
		t.Errorf("Expected car to be RED")
	}
}
