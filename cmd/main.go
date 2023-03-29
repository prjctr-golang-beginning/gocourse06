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
