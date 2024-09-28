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

func TestCheckCarColor(t *testing.T) {
	car := NewCar("RJ-14-HH-6294", RED)

	if !car.IsSameColor(RED) {
		t.Errorf("Expected car to be RED")
	}
}
