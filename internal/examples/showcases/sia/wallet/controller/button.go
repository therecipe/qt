package cwallet

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

var ButtonController *walletButtonController

type walletButtonController struct { //TODO: fix name clash
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (c *walletButtonController) init() {
	ButtonController = c

	c.ConnectClicked(cdialog.Controller.Show)
}
