package cleft

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type logoController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func() `signal:"clicked"`
}

func (c *logoController) init() {
	c.ConnectClicked(c.clicked)
}

func (c *logoController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://sia.tech", 0))
}
