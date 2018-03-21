package left

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { leftButtonTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ButtonTemplate") }

type leftButtonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked"`
}

func (t *leftButtonTemplate) init() {
	c := cleft.NewButtonController(nil)

	t.ConnectClicked(c.Clicked)
}
