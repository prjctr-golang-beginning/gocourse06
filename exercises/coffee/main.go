package main

import "fmt"

// інтерфейс кави
type Coffee interface {
	Type() string
}

// структура арабіки
type Arabica struct{}

func (a *Arabica) Type() string {
	return "Arabica"
}

// структура бленду
type Blend struct{}

func (b *Blend) Type() string {
	return "Blend"
}

// інтерфейс кавоварки
type CoffeeMaker interface {
	SetCoffee(coffee Coffee)
	Brew() string
}

var _ CoffeeMaker = &MyCoffeeMaker{}

// структура кавоварки з композицією кави
type MyCoffeeMaker struct {
	coffee Coffee
}

func (m *MyCoffeeMaker) SetCoffee(coffee Coffee) {
	m.coffee = coffee
}

func (m *MyCoffeeMaker) Brew() string {
	return fmt.Sprintf("Brewing %s coffee", m.coffee.Type())
}

func main() {
	// створюємо кавоварку
	var myMaker CoffeeMaker

	myMaker = &MyCoffeeMaker{}

	// створюємо каву
	arabica := &Arabica{}
	blend := &Blend{}

	// варимо арабіку
	myMaker.SetCoffee(arabica)
	fmt.Println(myMaker.Brew())

	// варимо бленд
	myMaker.SetCoffee(blend)
	fmt.Println(myMaker.Brew())
}
