package dialog

import (
	"github.com/therecipe/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { receiveTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "ReceiveTemplate") }

type receiveTemplate struct {
	dialogTemplate

	_ string `property:"address"`
}

func (t *receiveTemplate) show(cident string) {
	if cident == "receive" {
		t.SetAddress(controller.Controller.Receive())
	}
}
