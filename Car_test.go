package main

import "testing"

func TestCheckCreatedCarShouldMatchItself(t *testing.T) {
	car := newCar("KA-01-HH-1234", RED)

	if !car.isIdenticalCar("KA-01-HH-1234") {
		t.Errorf("Expected car to be created with registration number KA-01-HH-1234")
	}
}

func TestCheckCarColor(t *testing.T) {
	car := newCar("RJ-14-HH-6294", RED)

	if !car.isSameColor(RED) {
		t.Errorf("Expected car to be RED")
	}
}
