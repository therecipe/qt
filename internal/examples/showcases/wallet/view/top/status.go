package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/view/top/controller"
)

func init() { statusTemplate_QmlRegisterType2("TopTemplate", 1, 0, "StatusTemplate") }

type statusTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string `property:"balance,<-(this.c)"`
	_ string `property:"delta,<-(this.c)"`

	c *controller.StatusController
}

func (t *statusTemplate) init() {
	t.c = controller.NewStatusController(nil)
}
