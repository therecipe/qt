package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
	lcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func(string) `signal:"search,auto"`
}

func (c *searchController) search(name string) {
	lcontroller.LeftController.Clicked("files")
	controller.FilesController.Model().Filter.SetFilterFixedString(name)
}
