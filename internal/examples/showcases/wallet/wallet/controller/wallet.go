package controller

import (
	"encoding/json"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/wallet/model"
)

var Controller *WalletController

type WalletController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.WalletModel `property:"model"`

	_ func(ID string) `signal:"doubleClicked,auto"`
}

func (c *WalletController) init() {
	Controller = c

	c.SetModel(model.NewWalletModel(nil))

	go c.loop()
}

func (c *WalletController) doubleClicked(ID string) {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com", 0))
}

func (c *WalletController) loop() {
	for range time.NewTicker(1 * time.Second).C {
		if DEMO {
			var dt []model.Transaction
			json.Unmarshal([]byte(DEMO_TRANSACTIONS), &dt)
			c.Model().UpdateWith(dt)
		}
	}
}
