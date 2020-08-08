package controller

import "github.com/StarAurryon/qt/core"

var LeftController *leftController

type leftController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (c *leftController) init() {
	LeftController = c
}
