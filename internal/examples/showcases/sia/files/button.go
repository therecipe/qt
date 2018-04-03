package files

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
)

func init() { buttonTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (t *buttonTemplate) init() {
	c := controller.ButtonController
	if c == nil {
		c = controller.NewButtonController(nil)
	}

	t.ConnectClicked(c.Clicked)
}
