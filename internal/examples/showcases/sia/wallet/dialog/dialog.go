package dialog

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

type dialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(bool)          `signal:"blur"`
}

func (t *dialogTemplate) init() {
	c := controller.Controller
	if c == nil {
		c = controller.NewDialogController(nil)
	}

	c.ConnectShow(t.show)
	t.ConnectBlur(c.Blur)
}

func (t *dialogTemplate) show(cident string) {
	if controller.Controller.IsLocked() {
		t.Show("unlock")
	} else {
		t.Show(cident)
	}
}
