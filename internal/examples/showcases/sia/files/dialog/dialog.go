package dialog

import (
	"github.com/therecipe/qt/core"

	fcdialog "github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
	wcdialog "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

type filesDialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(bool)          `signal:"blur"`
}

func (t *filesDialogTemplate) init() {
	c := fcdialog.Controller
	if c == nil {
		c = fcdialog.NewFilesDialogController(nil)
	}

	c.ConnectShow(t.show)
	t.ConnectBlur(c.Blur)
}

func (t *filesDialogTemplate) show(cident string) {
	if fcdialog.Controller.IsLocked() {
		wcdialog.Controller.Show("unlock")
	} else {
		t.Show(cident)
	}
}
