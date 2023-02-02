package controller

import (
	"github.com/bluszcz/cutego/core"

	_ "github.com/bluszcz/cutego/internal/examples/showcases/wallet/files/dialog/controller"
)

var ButtonController *buttonController

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.Controller.Show)"`
}

func (c *buttonController) init() {
	ButtonController = c
}
