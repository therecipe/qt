package view

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

func init() { stackTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "StackTemplate") }

type stackTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(code string) `signal:"clicked"`
}

func (t *stackTemplate) init() {
	c := cview.NewStackController(nil)

	c.ConnectClicked(t.Clicked)
}
