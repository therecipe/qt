package main

import (
	"os"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/gui"
	"github.com/StarAurryon/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/qml/calc.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
