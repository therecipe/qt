package controller

import (
	"github.com/akiyosi/qt/core"

	"github.com/akiyosi/qt/internal/examples/showcases/wallet/controller"
	dcontroller "github.com/akiyosi/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() {
	if false {
		controller.Controller = nil
	}
}

type LockController struct {
	core.QObject

	_ bool `property:"locked,<-(controller.Controller)"`

	_ func() `signal:"change,auto"`
}

func (c *LockController) change() {
	if c.IsLocked() {
		dcontroller.Controller.Show("unlock")
	} else {
	}
}
