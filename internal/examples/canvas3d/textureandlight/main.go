//source: http://doc.qt.io/qt-5/qtcanvas3d-textureandlight-example.html

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

	//Turn on all the logging categories of Canvas3D
	var loggingFilter = `qt.canvas3d.info.debug=true
	   qt.canvas3d.rendering.debug=true
	   qt.canvas3d.rendering.warning=true
	   qt.canvas3d.glerrors.debug=true`
	core.QLoggingCategory_SetFilterRules(loggingFilter)

	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var viewer = quick.NewQQuickView(nil)

	var extraImportPath string
	if runtime.GOOS == "windows" {
		extraImportPath = "%v/../../../../%v"
	} else {
		extraImportPath = "%v/../../../%v"
	}

	viewer.Engine().AddImportPath(fmt.Sprintf(extraImportPath, app.ApplicationDirPath(), "qml"))
	viewer.SetSource(core.NewQUrl3("qrc:/main.qml", 0))

	viewer.SetTitle("Textured and Lit Cube")
	viewer.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	viewer.SetColor(gui.NewQColor6("#fafafa"))
	viewer.Show()

	app.Exec()
}
