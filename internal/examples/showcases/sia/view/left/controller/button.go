package cleft

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(code string) `signal:"clicked"`
}

func (c *buttonController) init() {
	//lazy binding to the view/stack controller
	c.ConnectClicked(func(code string) { cview.StackController.Clicked(code) })
}
