// +build !qml

package main

import (
	"os"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/widgets"

	"github.com/bluszcz/cutego/internal/examples/sql/masterdetail_qml/controller"

	"github.com/bluszcz/cutego/internal/examples/sql/masterdetail_qml/view"
)

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	view.NewViewController(nil, 0).Show()

	qApp.Exec()
}
