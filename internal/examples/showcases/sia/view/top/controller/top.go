package controller

import "github.com/therecipe/qt/core"

type topController struct {
	core.QObject

	_ func() `constructor:"init"`
}

func (c *topController) init() {

}
