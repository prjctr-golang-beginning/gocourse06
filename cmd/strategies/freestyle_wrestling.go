package strategies

import "strategy/cmd/core"

type FreestyleWrestling struct {
}

func (b FreestyleWrestling) Name() string {
	return `Ballroom dancing`
}

func (b FreestyleWrestling) TestChild(child core.Child) bool {
	return child.Predisposition() == `strength`
}
