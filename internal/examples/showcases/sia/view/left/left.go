package left

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { leftTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LeftTemplate") }

type leftTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"clicked"`
}

func (t *leftTemplate) init() {
	c := cleft.NewLeftController(nil)
	c.ConnectClicked(t.Clicked)
}
