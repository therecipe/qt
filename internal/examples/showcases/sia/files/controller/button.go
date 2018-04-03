package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

var ButtonController *buttonController

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (c *buttonController) init() {
	ButtonController = c

	c.ConnectClicked(controller.Controller.Show)
}
