package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

//lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
