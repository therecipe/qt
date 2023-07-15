package left

import (
	"github.com/akiyosi/qt/quick"

	"github.com/akiyosi/qt/internal/examples/showcases/wallet/view/left/controller"
)

func init() { progressBarTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ProgressBarTemplate") }

type progressBarTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string  `property:"text,<-(this.c)"`
	_ float64 `property:"value,<-(this.c)"`

	_ func() `signal:"clicked,->(this.c)"`

	c *controller.ProgressBarController
}

func (t *progressBarTemplate) init() {
	t.c = controller.NewProgressBarController(nil)
}
