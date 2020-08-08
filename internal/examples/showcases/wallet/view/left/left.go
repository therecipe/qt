package left

import (
	"github.com/StarAurryon/qt/quick"

	_ "github.com/StarAurryon/qt/internal/examples/showcases/wallet/view/left/controller"
)

func init() { leftTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LeftTemplate") }

type leftTemplate struct {
	quick.QQuickItem

	_ func(cident string) `signal:"Clicked,<-(controller.NewLeftController(nil))"`
}
