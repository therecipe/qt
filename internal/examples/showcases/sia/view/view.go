package view

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

func init() { viewTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "ViewTemplate") }

type viewTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(b bool) `signal:"blur"`
}

func (t *viewTemplate) init() {
	c := controller.NewViewController(nil)

	c.ConnectBlur(t.Blur)
}
