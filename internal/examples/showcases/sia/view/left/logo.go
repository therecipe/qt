package left

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { logoTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LogoTemplate") }

type logoTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func() `signal:"clicked"`
}

func (t *logoTemplate) init() {
	c := controller.NewLogoController(nil)

	t.ConnectClicked(c.Clicked)
}
