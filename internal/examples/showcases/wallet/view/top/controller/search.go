package controller

import (
	"github.com/StarAurryon/qt/core"

	"github.com/StarAurryon/qt/internal/examples/showcases/wallet/files/controller"
	lcontroller "github.com/StarAurryon/qt/internal/examples/showcases/wallet/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func(string) `signal:"search,auto"`
}

func (c *searchController) search(name string) {
	lcontroller.LeftController.Clicked("files")
	controller.FilesController.Model().Filter.SetFilterFixedString(name)
}
