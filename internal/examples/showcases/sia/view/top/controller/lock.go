package ctop

import (
	"net/http"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

type lockController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ bool `property:"locked"`

	_ func() `signal:"change"`
}

func (c *lockController) init() {

	c.ConnectIsLocked(controller.Controller.IsLocked)
	c.ConnectSetLocked(controller.Controller.SetLocked)
	controller.Controller.ConnectLockedChanged(c.LockedChanged)

	c.ConnectChange(c.change)
}

func (c *lockController) change() {
	if c.IsLocked() {
		cdialog.Controller.Show("unlock")
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
