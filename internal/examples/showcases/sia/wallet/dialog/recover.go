package dialog

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { recoverTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "RecoverTemplate") }

type recoverTemplate struct {
	dialogTemplate

	_ func() `constructor:"init"`

	_ func(string) *core.QVariant `slot:"recover"`
}

func (t *recoverTemplate) init() {
	t.ConnectRecover(cdialog.Controller.Recover)
}
