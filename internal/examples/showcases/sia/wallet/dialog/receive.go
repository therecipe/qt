package dialog

import (
	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { receiveTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "ReceiveTemplate") }

type receiveTemplate struct {
	dialogTemplate

	_ func() `constructor:"init"`

	_ string `property:"address"`
}

func (t *receiveTemplate) init() {
	t.ConnectShow(func(cident string) {
		if cident == "receive" {
			t.SetAddress(cdialog.Controller.Receive())
		}
	})
}
