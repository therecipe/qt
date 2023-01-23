package controller

import (
	"github.com/bluszcz/cutego/core"

	"github.com/bluszcz/cutego/internal/examples/showcases/wallet/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

//lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
