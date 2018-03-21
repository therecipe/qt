package dialog

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { sendTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "SendTemplate") }

type sendTemplate struct {
	dialogTemplate

	_ func() `constructor:"init"`

	_ func(string, string) *core.QVariant `slot:"send"`
}

func (t *sendTemplate) init() {
	t.ConnectSend(cdialog.Controller.Send)
}
