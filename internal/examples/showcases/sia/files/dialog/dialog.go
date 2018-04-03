package dialog

import (
	"github.com/therecipe/qt/core"

	fcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
	wcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

type filesDialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(bool)          `signal:"blur"`
}

func (t *filesDialogTemplate) init() {
	c := fcontroller.Controller
	if c == nil {
		c = fcontroller.NewDialogController(nil)
	}

	c.ConnectShow(t.show)
	t.ConnectBlur(c.Blur)
}

func (t *filesDialogTemplate) show(cident string) {
	if fcontroller.Controller.IsLocked() {
		wcontroller.Controller.Show("unlock")
	} else {
		t.Show(cident)
	}
}
