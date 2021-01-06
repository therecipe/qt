package files

import (
	"github.com/dev-drprasad/qt/quick"

	"github.com/dev-drprasad/qt/internal/examples/showcases/wallet/files/controller"
)

func init() { buttonTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.ButtonController)"`
}

func (t *buttonTemplate) init() {
	if controller.ButtonController == nil {
		controller.NewButtonController(nil)
	}
}
