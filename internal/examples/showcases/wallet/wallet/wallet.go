package wallet

import (
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/quick"

	"github.com/bluszcz/cutego/internal/examples/showcases/wallet/wallet/controller"
)

func init() { walletTemplate_QmlRegisterType2("WalletTemplate", 1, 0, "WalletTemplate") }

type walletTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractTableModel `property:"WalletModel"`

	_ func(string) `signal:"doubleClicked,->(this.c)"`

	c *controller.WalletController
}

func (t *walletTemplate) init() {
	t.c = controller.NewWalletController(nil)

	t.SetWalletModel(t.c.Model())
}
