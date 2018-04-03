package left

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { buttonTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (t *buttonTemplate) init() {
	c := controller.NewButtonController(nil)

	t.ConnectClicked(c.Clicked)
}
