package strategies

import "strategy/cmd/core"

type Swimming struct {
}

func (b Swimming) Name() string {
	return `Swimming`
}

func (b Swimming) TestChild(child core.Child) bool {
	return child.Predisposition() == `brain`
}
