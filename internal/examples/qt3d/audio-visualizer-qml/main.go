//source: https://doc.qt.io/qt-5/qt3d-audio-visualizer-qml-example.html

package main

import (
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var format = gui.NewQSurfaceFormat()
	if true /*TODO: gui.QOpenGLContext__OpenGLModuleType == gui.QOpenGLContext__LibGL*/ {
		format.SetVersion(3, 2)
		format.SetProfile(gui.QSurfaceFormat__CoreProfile)
	}
	format.SetDepthBufferSize(24)
	format.SetStencilBufferSize(8)

	var view = quick.NewQQuickView(nil)
	view.SetFormat(format)
	view.Create()

	var touchSettings = NewTouchSettings(nil)
	view.RootContext().SetContextProperty("touchSettings", touchSettings)

	view.SetSource(core.NewQUrl3("qrc:/main.qml", 0))

	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetMaximumSize(core.NewQSize2(1820, 1080))
	view.SetMinimumSize(core.NewQSize2(300, 150))
	view.Show()

	app.Exec()
}
