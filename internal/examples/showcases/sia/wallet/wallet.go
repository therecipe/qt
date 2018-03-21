package wallet

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/controller"
)

func init() { walletTemplate_QmlRegisterType2("WalletTemplate", 1, 0, "WalletTemplate") }

type walletTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractTableModel `property:"WalletModel"`

	_ func(string) `signal:"doubleClicked"`
}

func (t *walletTemplate) init() {
	c := cwallet.NewWalletController(nil)

	t.SetWalletModel(c.Model())

	t.ConnectDoubleClicked(c.DoubleClicked)
}
