package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/view/top/controller"
)

func init() { lockTemplate_QmlRegisterType2("TopTemplate", 1, 0, "LockTemplate") }

type lockTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ bool `property:"locked,<-(this.c)"`

	_ func() `signal:"change,->(this.c)"`

	c *controller.LockController
}

func (t *lockTemplate) init() {
	t.c = controller.NewLockController(nil)
}
