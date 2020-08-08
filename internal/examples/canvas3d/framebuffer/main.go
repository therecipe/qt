//source: https://doc.qt.io/qt-5.11/qtcanvas3d-framebuffer-example.html

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

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

	viewer.SetTitle("Render into FrameBuffer")
	viewer.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	viewer.SetColor(gui.NewQColor6("#f2f2f2"))
	viewer.Show()

	app.Exec()
}
