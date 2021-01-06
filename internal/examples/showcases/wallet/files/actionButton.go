package files

import (
	"github.com/dev-drprasad/qt/quick"

	"github.com/dev-drprasad/qt/internal/examples/showcases/wallet/files/controller"
)

func init() { actionButtonTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "ActionButtonTemplate") }

type actionButtonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload,->(controller.ActionButtonController)"`
	_ func(string) `signal:"deleteRequest,->(controller.ActionButtonController)"`
}

func (t *actionButtonTemplate) init() {
	if controller.ActionButtonController == nil {
		controller.NewActionButtonController(nil)
	}
}
