package controller

import (
	"github.com/akiyosi/qt/core"

	_ "github.com/akiyosi/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
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
