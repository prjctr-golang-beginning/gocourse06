package main

import "fmt"

type CoffeeMaker interface {
	Brew() string
}

var _ CoffeeMaker = &DripCoffeeMaker{}
var _ CoffeeMaker = &DripCoffeeMaker{}

type DripCoffeeMaker struct{}

func (d *DripCoffeeMaker) Brew() string {
	return "Drip coffee"
}

type EspressoMachine struct{}

func (e *EspressoMachine) Brew() string {
	return "Espresso"
}

func main() {
	dripCoffeeMaker := &DripCoffeeMaker{}
	fmt.Println(dripCoffeeMaker.Brew())

	espressoMachine := &EspressoMachine{}
	fmt.Println(espressoMachine.Brew())
}
