package main

import (
	"strategy/cmd/core"
	"strategy/cmd/kindergarten"
	"strategy/cmd/strategies"
	"time"
)

func main() {
	// strategies set
	sportSet := []core.Sport{
		strategies.BallroomDancing{},
		strategies.Boxing{},
		strategies.FreestyleWrestling{},
		strategies.Swimming{},
	}
	// children generator
	CheckChildren(sportSet)

	// exercises
	// 1. Створити інтерфейс кавоварки і його імплементацію
	// 2. Навчити кавоварку варити на арабіці і на бленді (інтерфейс кави і дви стратегії)
}

func CheckChildren(sportSet []core.Sport) {
	for {
		ch := kindergarten.GenerateChildren()
		for _, sport := range sportSet {
			// give sport to the child
			if sport.TestChild(&ch) {
				ch.Learn(sport)
			}
		}
		time.Sleep(time.Second)
	}
}
