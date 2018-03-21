package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { colorTemplate_QmlRegisterType2("TopTemplate", 1, 0, "ColorTemplate") }

type colorTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func() `signal:"change"`
}

func (t *colorTemplate) init() {
	c := ctop.NewColorController(nil)

	t.ConnectChange(c.Change)
}
