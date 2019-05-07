package dialog

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

type dialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show,<-(controller.Controller)"`
	_ func(bool)          `signal:"blur,->(controller.Controller)"`
}

func (t *dialogTemplate) init() {
	if controller.Controller == nil {
		controller.NewDialogController(nil)
	}
}

func (t *dialogTemplate) show(cident string) {
	if controller.Controller.IsLocked() {
		t.Show("unlock")
	} else {
		t.Show(cident)
	}
}
