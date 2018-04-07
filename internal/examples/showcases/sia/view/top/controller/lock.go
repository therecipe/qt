package controller

import (
	"net/http"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	dcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

type LockController struct {
	core.QObject

	_ bool `property:"locked,<-(controller.Controller)"`

	_ func() `signal:"change,auto"`
}

func (c *LockController) change() {
	if c.IsLocked() {
		dcontroller.Controller.Show("unlock")
	} else {
		req, _ := controller.Client.NewRequest("POST", "/wallet/lock", nil) //TODO:
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			println(err.Error())
		} else {
			c.SetLocked(true)
			res.Body.Close()
		}
	}
}
