package controller

import (
	"github.com/therecipe/qt/core"

	maincontroller "github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/files/dialog/controller"
)

var ActionButtonController *actionButtonController

type actionButtonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload,->(controller.Controller)"`
	_ func(string) `signal:"deleteRequest,auto"`
}

func (c *actionButtonController) init() {
	ActionButtonController = c
}

func (c *actionButtonController) deleteRequest(name string) {
	err := maincontroller.Client.RenterDeletePost(name)
	if err != nil {
		println("Couldn't delete file:", err.Error())
	}
}
