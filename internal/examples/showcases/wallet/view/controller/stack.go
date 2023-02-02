package controller

import "github.com/bluszcz/cutego/core"

var StackController *stackController

type stackController struct {
	core.QObject

	_ func() `constructor:"init`

	_ func(string) `signal:"clicked"`
}

func (c *stackController) init() {
	StackController = c
}
