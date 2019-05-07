package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/controller"
)

type StatusController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"balance"`
	_ string `property:"delta"`

	_ func(interface{}) `signal:"wallet,<-(controller.Controller.WalletChanged)"`
}

func (c *StatusController) init() {
	if controller.DEMO {
		c.wallet(nil)
	}
}

func (c *StatusController) wallet(wgI interface{}) {
	if controller.DEMO {
		c.SetBalance("12345 C")
		c.SetDelta("67890 C")
	}
}
