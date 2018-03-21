package wallet

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/controller"
)

func init() { walletButtonTemplate_QmlRegisterType2("WalletTemplate", 1, 0, "ButtonTemplate") }

type walletButtonTemplate struct { //TODO: fix name clash
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (t *walletButtonTemplate) init() {
	c := cwallet.ButtonController
	if c == nil {
		c = cwallet.NewWalletButtonController(nil)
	}

	t.ConnectClicked(c.Clicked)
}
