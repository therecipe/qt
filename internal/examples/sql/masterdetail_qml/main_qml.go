// +build qml

package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/controller"

	_ "github.com/therecipe/qt/internal/examples/sql/masterdetail_qml/view"
)

const PRODUCTION = false

func init() {
	if !PRODUCTION {
		os.Setenv("QML_DISABLE_DISK_CACHE", "true")
	}
}

func main() {
	var path string
	if PRODUCTION {
		path = "qrc:/qml/view.qml"
	} else {
		path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", "internal", "examples", "sql", "masterdetail_qml", "view", "qml", "view.qml")
	}

	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	app := qml.NewQQmlApplicationEngine(nil)
	if PRODUCTION {
		app.AddImportPath("qrc:/qml/")
	} else {
		app.AddImportPath("./view/detail/qml")
		app.AddImportPath("./view/album/qml")
		app.AddImportPath("./view/artist/qml")
		app.AddImportPath("./view/dialog/qml")
	}
	app.Load(core.NewQUrl3(path, 0))

	widgets.QApplication_Exec()
}
