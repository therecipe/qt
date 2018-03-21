package files

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
)

func init() { filesButtonTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "ButtonTemplate") }

type filesButtonTemplate struct { //TODO: fix name clash
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (t *filesButtonTemplate) init() {
	c := cfiles.ButtonController
	if c == nil {
		c = cfiles.NewFilesButtonController(nil)
	}

	t.ConnectClicked(c.Clicked)
}
