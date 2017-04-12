//source: http://doc.qt.io/qt-5/qtcanvas3d-threejs-planets-example.html

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {

	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var viewer = quick.NewQQuickView(nil)

	var extraImportPath string
	if runtime.GOOS == "windows" {
		extraImportPath = "%v/../../../../%v"
	} else {
		extraImportPath = "%v/../../../%v"
	}

	viewer.Engine().AddImportPath(fmt.Sprintf(extraImportPath, app.ApplicationDirPath(), "qml"))
	viewer.SetSource(core.NewQUrl3("qrc:/planets.qml", 0))

	viewer.SetTitle("Qt Canvas 3D Examples - Planets")
	viewer.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	viewer.SetColor(gui.NewQColor2(core.Qt__black))
	viewer.Show()

	app.Exec()
}
