package cfiles

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

var ActionButtonController *actionButtonController

type actionButtonController struct { //TODO: fix name clash
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload"`
	_ func(string) `signal:"deleteRequest"`
}

func (c *actionButtonController) init() {
	ActionButtonController = c

	c.ConnectShowDownload(cdialog.Controller.ShowDownload)
	c.ConnectDeleteRequest(c.deleteRequest)
}

func (c *actionButtonController) deleteRequest(name string) {
	err := controller.Client.RenterDeletePost(name)
	if err != nil {
		println("Couldn't delete file:", err.Error())
	}
}
