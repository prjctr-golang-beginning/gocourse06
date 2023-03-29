package strategies

import "strategy/cmd/core"

type Swimming struct {
}

func (b Swimming) Name() string {
	return `Ballroom dancing`
}

func (b Swimming) TestChild(child core.Child) bool {
	return child.Predisposition() == `brain`
}
