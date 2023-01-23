package controller

import (
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
)

type logoController struct {
	core.QObject

	_ func() `signal:"clicked,auto"`
}

func (c *logoController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com/", 0))
}
