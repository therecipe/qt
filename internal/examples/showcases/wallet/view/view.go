package view

import (
	"github.com/bluszcz/cutego/quick"

	_ "github.com/bluszcz/cutego/internal/examples/showcases/wallet/view/controller"
)

func init() { viewTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "ViewTemplate") }

type viewTemplate struct {
	quick.QQuickItem

	_ func(b bool) `signal:"Blur,<-(controller.NewViewController(nil))"`
}
