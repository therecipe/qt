package ctop

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/files/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"search"`
}

func (c *searchController) init() {
	c.ConnectSearch(c.search)
}

func (c *searchController) search(name string) {
	cleft.LeftController.Clicked("files")
	cfiles.FilesController.Model().Filter.SetFilterFixedString(name)
}
