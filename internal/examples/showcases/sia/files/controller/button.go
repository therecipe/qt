package cfiles

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

var ButtonController *filesButtonController

type filesButtonController struct { //TODO: fix name clash
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (c *filesButtonController) init() {
	ButtonController = c

	c.ConnectClicked(cdialog.Controller.Show)
}
