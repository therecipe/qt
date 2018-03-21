package ctop

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/theme/controller"
)

type colorController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func() `signal:"change"`
}

func (c *colorController) init() {
	//lazy binding to the (qml singleton) theme controller
	c.ConnectChange(func() { ctheme.Controller.Change() })
}
