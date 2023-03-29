package strategies

import "strategy/cmd/core"

type Boxing struct {
}

func (b Boxing) Name() string {
	return `boxing`
}

func (b Boxing) TestChild(child core.Child) bool {
	return child.Predisposition() == `nothing`
}
